package jvvt

import (
	"testing"
)

func TestEncodeComponent(t *testing.T) {
	var byteData []byte
	newStr := "abc123!?$*&()'-=@~"
	byteData = []byte(newStr)

	result := encodeComponent(byteData)
	if result != "YWJjMTIzIT8kKiYoKSctPUB-" {
		t.Errorf("returned = %s, expected = \"YWJjMTIzIT8kKiYoKSctPUB-\"", result)
	}

}

func TestDecodeComponent(t *testing.T) {
	stringData := "YWJjMTIzIT8kKiYoKSctPUB-"

	result := decodeComponent(stringData)
	if string(result) != "abc123!?$*&()'-=@~" {
		t.Errorf("returned = %s, expected = abc123!?$*&()'-=@~", string(result))
	}
}
