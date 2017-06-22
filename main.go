package main

import (
	"encoding/json"
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("You must provide a token")
		return
	}

	token, err := jwt.Parse(args[0], func(token *jwt.Token) (interface{}, error) {
		return []byte(nil), nil
	})

	if token == nil {
		fmt.Println(err)
	} else {
		b, _ := json.MarshalIndent(token.Claims, "", " ")
		fmt.Println(string(b))
	}
}
