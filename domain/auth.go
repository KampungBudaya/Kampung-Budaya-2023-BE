package domain

import "strings"

type AuthRequest struct {
	Token string `json:"token"`
}

type AuthResponse struct {
	SignedToken string   `json:"signedToken"`
	Panels      []string `json:"panels"`
}

func (res *AuthResponse) Populate(rolesArg string) {
	roles := strings.Split(rolesArg, ", ")

	for _, role := range roles {
		switch role {
		case "Super Admin":
			res.Panels = []string{"Participants"}
			return

		case "Acara":
			res.Panels = append(res.Panels, "Venue")
		}
	}
}
