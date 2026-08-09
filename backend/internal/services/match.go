package services

import (
	db "backend/database"
	dto "backend/internal/dtos"
	"backend/pkg"
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func CreateMatch(req dto.CreateMatchRequest, hostID string) (*dto.CreateMatchResponse, error) {
	matchID := pkg.GenerateRandomString(7)
	matchState := fiber.Map{
		"match_id":    matchID,
		"host_id":     hostID,
		"status":      "waiting",
		"difficulty":  req.Difficulty,
		"start_page":  "",
		"target_page": "",
		"players": fiber.Map{
			hostID: fiber.Map{
				"name":         req.DisplayName,
				"path_history": []string{},
				"clicks":       0,
			},
		},
	}
	cacheKey := "match:" + matchID
	cost := int64(1) // 1 cost = 1 time, total items we can store = 110
	saved := db.Cache.SetWithTTL(cacheKey, matchState, cost, time.Hour)
	db.Cache.Wait() // to ensure async op completes
	if !saved {
		log.Error("Failed to allocate cache space for match")
		return nil, errors.New("Failed to allocate cache space for match")
	}
	return &dto.CreateMatchResponse{
		MatchID:   matchID,
		InviteURL: "/match/" + matchID,
	}, nil
}

func JoinMatch(req dto.JoinMatchRequest, clientID string) error {
	matchID := req.MatchID
	matchLobbyRaw, found := db.Cache.Get(matchID)
	if !found {
		return fiber.ErrNotFound
	}
	matchLobby := matchLobbyRaw.(fiber.Map)
	players := matchLobby["players"].(fiber.Map)
	if len(players) >= 2 {
		return fiber.ErrConflict
	}
	// just a check to ensure, same user isnt trying to connect again
	_, exists := players[clientID]
	if exists {
		return fiber.ErrBadRequest
	}
	players[clientID] = fiber.Map{
		"name":         req.DisplayName,
		"path_history": []string{},
		"clicks":       0,
	}
	matchLobby["players"] = players
	matchLobby["status"] = "generating"
	db.Cache.Set(matchID, matchLobby, 1)
	db.Cache.Wait()
}
