package jvvt

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"hash"
	"strings"
)

// Claims : These are claims that have specific meanings attached
// to them

type Header struct {
	Algorithm string `json:"alg"`
	Toketype  string `json:"typ"`
}

type JvvtObj struct {
	signingHash hash.Hash
	Head        Header
}

/*

func Verify(rawToken string) bool {

}*/

func NewJVVT(secret string) JvvtObj {
	return JvvtObj{
		signingHash: hmac.New(sha256.New, []byte(secret)),
		Head: Header{
			Algorithm: "HS256",
			Toketype:  "jwt",
		},
	}
}

func (j *JvvtObj) signToken(tokenUnsigned string) []byte {
	j.signingHash.Write([]byte(tokenUnsigned))
	sha := j.signingHash.Sum(nil)
	j.signingHash.Reset()
	return sha
}

func (j *JvvtObj) GenerateToken(claims Claims) (string, error) {
	header, err := json.Marshal(j.Head)
	if err != nil {
		return "", err
	}
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	headerBase64 := encodeComponent(header)
	payloadBase64 := encodeComponent(payload)

	unSignedPart := headerBase64 + "." + payloadBase64
	//fmt.Println(unSignedPart)
	sha := j.signToken(unSignedPart)
	sign := encodeComponent([]byte(sha))

	return headerBase64 + "." + payloadBase64 + "." + sign, nil
}

func (j *JvvtObj) GetClaims(token string) (Claims, error) {
	tokenComps := strings.Split(token, ".")

	if len(tokenComps) != 3 {
		if len(tokenComps) == 2 {
			return Claims{}, errors.New("token cannot be unsigned")
		} else {
			return Claims{}, errors.New("token not in proper format")
		}
	}
	if !j.VerifySignature(token) {
		return Claims{}, errors.New("token not valid")
	}

	b64Payload := tokenComps[1]

	payload := decodeComponent(b64Payload)

	claims := NewClaims()
	if err := json.Unmarshal(payload, &claims); err != nil {
		return claims, errors.New("token data could not be converted to Claims obj: " + err.Error())
	}

	return claims, nil
}

func (j *JvvtObj) Verify(token string) (bool, error) {
	tokenComps := strings.Split(token, ".")

	if len(tokenComps) != 3 {
		return false, errors.New("token should consist of three components: header.payload.signature")
	}

	if !j.VerifySignature(token) {
		return false, nil
	}

	claims, err := j.GetClaims(token)

	if err != nil {
		return false, errors.New("error at GetClaims" + err.Error())
	}
	if claims.IsTokenExpried() {
		return false, nil
	}

	return true, nil
}

func (j *JvvtObj) VerifySignature(token string) bool {
	splitedStr := strings.Split(token, ".")

	b64Header := splitedStr[0]
	b64Payload := splitedStr[1]
	b64Sign := splitedStr[2]

	unSignedPart := b64Header + "." + b64Payload

	//fmt.Println(unSignedPart)
	generatedSign := j.signToken(unSignedPart)

	sign := encodeComponent([]byte(generatedSign))

	//fmt.Println(sign)

	return hmac.Equal([]byte(b64Sign), []byte(sign))
}

func encodeComponent(data []byte) string {
	b64urlSafe := b64.RawURLEncoding.EncodeToString(data)
	return b64urlSafe
}

func decodeComponent(data string) []byte {
	bytedata, _ := b64.RawURLEncoding.DecodeString(data)
	return bytedata
}
