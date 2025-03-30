package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ref > https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
const HashCost = bcrypt.DefaultCost

func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, HashCost)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return "", err
	}
	return string(hash), nil
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			log.Println("Invalid password")
		} else {
			log.Printf("Error comparing password: %v", err)
		}
		return false
	}
	return true
}
