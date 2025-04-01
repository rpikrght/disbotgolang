package token

import (
	"log"
	"os"
)

func String() (string, error) {
	tokenstr, err := os.ReadFile("./token/token.txt")
	if err != nil {
		log.Printf("opentokenerr: %v", err)
		return "", err
	}
	return string(tokenstr), nil
}
