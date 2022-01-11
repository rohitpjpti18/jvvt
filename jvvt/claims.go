package jvvt

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

func NewClaims() Claims {
	return Claims{}
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
