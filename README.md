# jvvt
library for implementation of json web tokens in go.

## Getting started

Import the jvvt package in your code
JvvtObj acts as a container for handling jwt tokens

#### Create a new JvvtObj
``` Go
    secretKey := "someLargeSeceretKey" // ! IMPORTANT keep this key secret , 
    jvvtObj := jvvt.NewJvvtObj(secretKey)
``` 

claims is a struct type which acts as a payload for jwt token
``` Go
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
```

#### creating a new claim object
``` Go
myclaims := jvvt.NewClaims()
myclaims.Issuer = "dumbledore"
myclaims.Subject = "entry.to.hogwarts"
myclaims.IssuedAt = time.Now().Unix()                   // IssuedAt takes unix time only 
myclaims.Expiration = time.Now().AddDate(0,0,2).Unix()  // Expiration date of token 
```

#### generating a new token
``` Go
token := jvvt.GenerateToken(claims)
```

#### verifying the token 
``` Go
isValid, err := jvvt.Verify(token)
```

#### getting claims from the raw token
``` Go
claims, err := jvvt.GetClaims(token)
```
