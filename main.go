package main

import (
	"fmt"
	"jvvt-main/jvvt"
)

func main() {
	var secret string = "somesecret"
	var jvvtObj = jvvt.NewJVVT(secret)

	claims := jvvt.NewClaims()

	claims.Issuer = "test-app"

	someToken := jvvtObj.GenerateToken(claims)
	fmt.Println(someToken)

	fmt.Println(jvvtObj.Verify(someToken))

	/* 	algorithm := jwt.HmacSha256("ThisIsTheSecret")
	   	claims := jwt.NewClaim()
	   	claims.Set("Role", "Admin")
	   	token, err := algorithm.Encode(claims)
	   	if err != nil {
	   		panic(err)
	   	}

	   	fmt.Println(token) */
}
