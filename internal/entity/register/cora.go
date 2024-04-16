package entity_cora_api

type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	BusinessID       string `json:"business_id"`
	ExpiresIn        int    `json:"expires_in"`
	NotBeforePolicy  int    `json:"not-before-policy"`
	PersonID         string `json:"person_id"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	Scope            string `json:"scope"`
	TokenType        string `json:"token_type"`
}
