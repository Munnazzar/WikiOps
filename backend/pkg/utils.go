package pkg

import (
	"math/rand/v2"
	"net/http"
	"strings"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(n int) string {
	b := make([]byte, n)

	for i := 0; i < n; {
		r := rand.Uint64()

		for j := 0; j < 12 && i < n; j++ {
			b[i] = charset[r%36]
			r /= 36
			i++
		}
	}
	return string(b)
}

var wikiClient = &http.Client{}

func RandomWikipediaArticle() (string, error) {
	resp, err := wikiClient.Get("https://en.wikipedia.org/wiki/Special:Random")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return strings.TrimPrefix(resp.Request.URL.String(), "/wiki/"), nil
}
