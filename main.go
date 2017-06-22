package main

import (
	"encoding/json"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	token, err := jwt.Parse(os.Args[1:][0], func(token *jwt.Token) (interface{}, error) {
		return []byte(nil), nil
	})

	if token == nil {
		fmt.Println(err)
	} else {
		b, _ := json.MarshalIndent(token.Claims, "", " ")
		fmt.Println(string(b))
	}
}
