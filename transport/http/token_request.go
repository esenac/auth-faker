package http

// TokenRequest represents a request for a new token
type TokenRequest struct {
	CustomClaims map[string]interface{} `json:"custom_claims"`
	Subject      string                 `json:"subject"`
	Audience     string                 `json:"audience"`
	Scope        string                 `json:"scope"`
	Issuer       string                 `json:"issuer"`
}
