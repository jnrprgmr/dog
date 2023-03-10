package handlers

import (
	"net/http"
)

func (h *Handlers) TwitchAuthHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	h.twitch.Code = code
	resp, err := h.twitch.Client.RequestUserAccessToken(h.twitch.Code)
	if err != nil {
		// handle error
	}

	// Set the access token on the client
	h.twitch.Token = resp.Data.AccessToken
	h.twitch.RefreshToken = resp.Data.RefreshToken
	h.twitch.Client.SetUserAccessToken(h.twitch.Token)
	http.Redirect(w, r, "http://localhost:8080/twitch", http.StatusSeeOther)
}
