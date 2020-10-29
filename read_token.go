package jvvt

import (
	"strings"
)

func splitToken(str string) []string{
	token := strings.Split(str, ".");

	return token;
}


