package jvvt

import "time"

// payload data for jwt token
type Claims struct {
	Issuer       string                 `json:"iss,omitempty"`
	Subject      string                 `json:"sub,omitempty"`
	Audience     string                 `json:"aud,omitempty"`
	Expiration   int64                  `json:"exp,omitempty"`
	NotBefore    int64                  `json:"nbf,omitempty"`
	IssuedAt     int64                  `json:"iat,omitempty"`
	OtherDetails map[string]interface{} `json:"other,omitempty"`
	JWTID        string                 `json:"jti,omitempty"`
}

// get new claims object
func NewClaims() Claims {
	return Claims{}
}

// check if token is expired or not
func (c *Claims) IsTokenExpried() bool {
	exp := time.Unix(c.Expiration, 0)
	return exp.Before(time.Now())
}

/*
func (rc *RegisteredClaims) Valid() error {
	vErr := new(ValidationError)
	now := time.Now().Unix()

	if rc.VerifyExpiration(now) == false {
		vErr.ErrorCode = 2;
	}
}
*/
