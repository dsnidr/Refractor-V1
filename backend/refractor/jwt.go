package refractor

// TokenPair represents a pair of JWTs
type TokenPair struct {
	AuthToken    string `json:"authToken"`
	RefreshToken string `json:"refreshToken"`
}
