package auth

import (
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func authenticate(token string) bool {
	_, err := jwt.DecodeSegment(token)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func GetTimesKey(token string) string {
	parts := strings.Split(token, ".")
	var decoded []byte
	if authenticate(token) {
		decoded, _ := jwt.DecodeSegment(parts[1])
		fmt.Println(decoded)
	}
	return string(decoded)
}
