package jvvt

// RegisteredClaims : These are claims that have specific meanings attached
// to them
type RegisteredClaims struct {
	Issuer     string `json:"iss,omitempty"`
	Subject    string `json:"sub,omitempty"`
	Audience   string `json:"aud,omitempty"`
	Expiration int64  `json:"exp,omitempty"`
	NotBefore  int64  `json:"nbf,omitempty"`
	IssuedAt   int64  `json:"iat,omitempty"`
	JWTID      string `json:"jti,omitempty"`
}

