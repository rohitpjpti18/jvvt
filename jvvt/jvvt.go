package jvvt

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
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

func getHS256Hash(secret string) hash.Hash {
	return hmac.New(sha256.New, []byte(secret))
}

func (j *JvvtObj) signToken(tokenUnsigned string) []byte {
	j.signingHash.Write([]byte(tokenUnsigned))
	sha := j.signingHash.Sum(nil)
	j.signingHash.Reset()
	return sha
}

func (j *JvvtObj) GenerateToken(claims Claims) string {
	header, err1 := json.Marshal(j.Head)
	payload, err2 := json.Marshal(claims)
	if err1 != nil || err2 != nil {
		fmt.Println("Error Marshalling the head or payload " + err1.Error() + err2.Error())
	}

	headerBase64 := encodeComponent(header)
	payloadBase64 := encodeComponent(payload)

	unSignedPart := headerBase64 + "." + payloadBase64
	//fmt.Println(unSignedPart)
	sha := j.signToken(unSignedPart)
	sign := encodeComponent([]byte(sha))

	return headerBase64 + "." + payloadBase64 + "." + sign
}

func (j *JvvtObj) Verify(token string) bool {
	splitedStr := strings.Split(token, ".")

	b64Header := splitedStr[0]
	b64Payload := splitedStr[1]
	b64Sign := splitedStr[2]

	unSignedPart := b64Header + "." + b64Payload

	fmt.Println(unSignedPart)
	generatedSign := j.signToken(unSignedPart)

	sign := encodeComponent([]byte(generatedSign))

	//fmt.Println(sign)

	return hmac.Equal([]byte(b64Sign), []byte(sign))
}

/*
func GetPayload() Token {

}
*/

func encodeComponent(data []byte) string {
	b64urlSafe := b64.RawURLEncoding.EncodeToString(data)
	return b64urlSafe
}

func decodeComponent(data string) []byte {
	bytedata, _ := b64.RawURLEncoding.DecodeString(data)
	return bytedata
}
