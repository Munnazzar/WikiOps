package dto

type CreateMatchRequest struct {
	Difficulty  string `json:"difficulty" validate:"required,oneof=easy medium hard"`
	DisplayName string `json:"display_name" validate:"required,min=3,max=25"`
}
type CreateMatchResponse struct {
	MatchID   string `json:"match_id"`
	InviteURL string `json:"invite_url"`
}

type JoinMatchRequest struct {
	MatchID     string `json:"match_id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,min=3,max=25"`
}
