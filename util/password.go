package util

import "golang.org/x/crypto/bcrypt"

func GenHashedPass(rawPass string) string {
	bPass := []byte(rawPass)
	hash, err := bcrypt.GenerateFromPassword(bPass, 6)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ValidatePass(rawPass string, hashStr string) bool {
	bPass := []byte(rawPass)
	bHash := []byte(hashStr)
	err := bcrypt.CompareHashAndPassword(bHash, bPass)
	if err != nil {
		return false
	}
	return true
}
