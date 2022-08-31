package util

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

// HashPassword
func HashPassword(s string) string {
	bs, err := bcrypt.GenerateFromPassword([]byte(s), MinCost)
	if err != nil {
		log.Fatalf("error when hash pw, %v", bs)
		return ""
	}
	bpass := string(bs)
	return bpass
}

func ComparePassword(hs string, ps string) bool {
	var isSuccess bool
	err := bcrypt.CompareHashAndPassword([]byte(hs), []byte(ps))
	if err == nil {
		isSuccess = true
		return isSuccess
	}
	return isSuccess
}
