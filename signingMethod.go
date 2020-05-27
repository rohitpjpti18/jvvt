package jvvt

import (
	"sync"
)

var signingMethods = map[String]func SigningMethod{}
var signingMethodLock = new(sync.RWMutex)

// SigningMethod is an interface that contains method for signing or verifying tokens
type SigningMethod interface{
	Verify(signingString, signature string, key interface{}) error
	Sign(signingString string, key interface{}) (string, error)
	Alg() string
}


