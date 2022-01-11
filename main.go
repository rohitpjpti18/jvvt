package main

import (
	"fmt"
	"jvvt-main/jvvt"
	"time"
)

func main() {
	var secret string = "somesecret"
	var jvvtObj = jvvt.NewJVVT(secret)

	claims := jvvt.NewClaims()

	claims.Issuer = "test-app"
	claims.Subject = "test-client"
	claims.IssuedAt = time.Now().Unix()                    // get current unix time
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix() // set expiration 48 hours from the time of issue

	someToken := jvvtObj.GenerateToken(claims)
	fmt.Println(someToken)

	fmt.Println(jvvtObj.VerifySignature(someToken))

	newClaim, err := jvvtObj.GetClaims(someToken)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(newClaim.Issuer)
	fmt.Println(newClaim.Subject)

	/* 	algorithm := jwt.HmacSha256("ThisIsTheSecret")
	   	claims := jwt.NewClaim()
	   	claims.Set("Role", "Admin")
	   	token, err := algorithm.Encode(claims)
	   	if err != nil {
	   		panic(err)
	   	}

	   	fmt.Println(token) */
}
