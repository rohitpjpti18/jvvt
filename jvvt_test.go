package jvvt

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

// test if byte array is properly being encoded
func TestEncodeComponent1(t *testing.T) {
	var byteData []byte
	newStr := "abc123!?$*&()'-=@~"
	byteData = []byte(newStr)

	result := encodeComponent(byteData)
	if result != "YWJjMTIzIT8kKiYoKSctPUB-" {
		t.Errorf("returned = %s, expected = \"YWJjMTIzIT8kKiYoKSctPUB-\"", result)
	}

}

// test if  is being decode into properly
func TestDecodeComponent(t *testing.T) {
	stringData := "YWJjMTIzIT8kKiYoKSctPUB-"

	result := decodeComponent(stringData)
	if string(result) != "abc123!?$*&()'-=@~" {
		t.Errorf("returned = %s, expected = abc123!?$*&()'-=@~", string(result))
	}
}

// test generation of token
func TestGenerateToken(t *testing.T) {
	jvvtTest := NewJVVT("somesecretkey")

	claims := NewClaims()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = "rohit.prajapati"
	claims.Subject = "user.login"
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix()

	token, err := jvvtTest.GenerateToken(claims)

	if err != nil {
		t.Errorf("token generation failed: " + err.Error())
	}
	if !jvvtTest.VerifySignature(token) {
		t.Errorf("token signature not valid")
	}

	decodedClaims, err := jvvtTest.GetClaims(token)
	if err != nil {
		t.Errorf(err.Error())
	}

	if decodedClaims.Expiration != claims.Expiration ||
		decodedClaims.IssuedAt != claims.IssuedAt ||
		decodedClaims.Issuer != claims.Issuer ||
		decodedClaims.Subject != claims.Subject {
		dcbyte, err := json.MarshalIndent(decodedClaims, "", "	")
		if err != nil {
			t.Errorf(err.Error())
		}
		cbyte, err := json.MarshalIndent(claims, "", "	")
		if err != nil {
			t.Errorf(err.Error())
		}
		t.Errorf("claims donot match expected: " + string(dcbyte) + " but got: " + string(cbyte))
	}
}

// test if right token is verified to be true
func TestVerifySignature1(t *testing.T) {
	jvvtTest := NewJVVT("somesecretkey")

	claims := NewClaims()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = "rohit.prajapati"
	claims.Subject = "user.login"
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix()

	token, err := jvvtTest.GenerateToken(claims)

	if err != nil {
		t.Errorf("token generation failed: " + err.Error())
	}
	if !jvvtTest.VerifySignature(token) {
		t.Errorf("token signature not valid")
	}
}

// test if corrupt token is verified to be false
func TestVerifySignature2(t *testing.T) {
	jvvtTest := NewJVVT("somesecretkey")

	claims := NewClaims()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = "rohit.prajapati"
	claims.Subject = "user.login"
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix()

	token, err := jvvtTest.GenerateToken(claims)

	if err != nil {
		t.Errorf("token generation failed: " + err.Error())
	}

	comps := strings.Split(token, ".")

	comps[0] = "sadkjsdafklejdskfa239-0dfsa" // corrupt header

	if jvvtTest.VerifySignature(comps[0] + "." + comps[1] + "." + comps[2]) {
		t.Errorf("signature not verified correctly")
	}
}

// test with corrupt token
func TestVerify1(t *testing.T) {
	jvvtTest := NewJVVT("somesecretkey")

	claims := NewClaims()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = "rohit.prajapati"
	claims.Subject = "user.login"
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix()

	token, err := jvvtTest.GenerateToken(claims)

	if err != nil {
		t.Errorf("token generation failed: " + err.Error())
	}

	comps := strings.Split(token, ".")

	comps[0] = "sadkjsdafklejdskfa239-0dfsa" // corrupt header

	if val, err := jvvtTest.Verify(comps[0] + "." + comps[1] + "." + comps[2]); val || err != nil {
		if err != nil {
			t.Errorf("Error occured during verification: " + err.Error())
		} else {
			t.Errorf("signature not verified correctly")
		}

	}
}

// test for expired token
func TestVerify2(t *testing.T) {
	jvvtTest := NewJVVT("somesecretkey")

	claims := NewClaims()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = "rohit.prajapati"
	claims.Subject = "user.login"
	claims.Expiration = time.Now().Unix()

	token, err := jvvtTest.GenerateToken(claims)

	time.Sleep(100 * time.Millisecond)

	if err != nil {
		t.Errorf("token generation failed: " + err.Error())
	}

	if val, err := jvvtTest.Verify(token); val || err != nil {
		if err != nil {
			t.Errorf("Error occured during verification: " + err.Error())
		} else {
			t.Errorf("signature not verified correctly")
		}

	}
}

// test with valid token
func TestVerify3(t *testing.T) {
	jvvtTest := NewJVVT("somesecretkey")

	claims := NewClaims()
	claims.IssuedAt = time.Now().Unix()
	claims.Issuer = "rohit.prajapati"
	claims.Subject = "user.login"
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix()

	token, err := jvvtTest.GenerateToken(claims)

	time.Sleep(100 * time.Millisecond)

	if err != nil {
		t.Errorf("token generation failed: " + err.Error())
	}

	if val, err := jvvtTest.Verify(token); !val || err != nil {
		if err != nil {
			t.Errorf("Error occured during verification: " + err.Error())
		} else {
			t.Errorf("signature not verified correctly")
		}

	}
}
