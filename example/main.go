package main

import (
	"fmt"
	"jvvt"
	"time"
)

func main() {
	var secret string = "somesecret"
	var jvvtObj = jvvt.NewJVVT(secret)

	claims := jvvt.NewClaims()

	claims.Issuer = "albus.dumbledore"
	claims.Subject = "permit.hogsmade.visit"
	claims.IssuedAt = time.Now().Unix()                    // get current unix time
	claims.Expiration = time.Now().AddDate(0, 0, 2).Unix() // set expiration 48 hours from the time of issue
	var otherdetails map[string]interface{} = make(map[string]interface{})
	otherdetails["username"] = "harry.potter"
	otherdetails["email"] = "harry.potter@hogwarts.com"

	claims.OtherDetails = otherdetails

	someToken, err := jvvtObj.GenerateToken(claims)
	if err != nil {
		fmt.Println(err.Error())
	}
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
