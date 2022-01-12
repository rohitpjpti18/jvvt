package jvvt

import (
	"testing"
	"time"
)

// test for expired tokens
func TestIsTokenExpired1(t *testing.T) {
	claims := NewClaims()

	claims.Issuer = "albus.dumbledore"
	claims.Subject = "owls"
	claims.IssuedAt = time.Now().Unix()
	claims.Expiration = time.Now().Unix()

	time.Sleep(100 * time.Millisecond)

	var isExp bool = claims.IsTokenExpried()

	if !isExp {
		t.Errorf("expected true from IsTokenExpired")
	}
}

// test for unexpired tokens
func TestIsTokenExpired2(t *testing.T) {
	claims := NewClaims()

	claims.Issuer = "albus.dumbledore"
	claims.Subject = "owls"
	claims.IssuedAt = time.Now().Unix()
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix()

	var isExp bool = claims.IsTokenExpried()

	if isExp {
		t.Errorf("expected false from IsTokenExpired")
	}
}
