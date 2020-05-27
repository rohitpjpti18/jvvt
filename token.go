package jvvt

type Token struct {
	Raw       string
	Method    SigningMethod
	Header    map[string]interface{}
	Claims    Claims
	Signature string
	Valid     bool
}

// New: Creates a new token. Taken a signing method
func New(method SigningMethod) *token {
	return NewWithClaims(method, MapClaims{})
}


func NewWithClaims(method SigningMethod, claims Claims) *Token {
	return &Token{
		Header: map[string]interface{}{
			"type": "JWT",
			"alg": method.Alg(),
		},
		Claims: claims,
		Method: method,
	}
}