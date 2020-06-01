package jvvt

import (
	b64 "encoding/base64"
)

func encodeComponent(data []byte) string {
	b64urlSafe := b64.URLEncoding.EncodeToString(data)
	return b64urlSafe
}

func decodeComponent(data string) []byte {
	bytedata, _ := b64.URLEncoding.DecodeString(data)
	return bytedata
}
