package security

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ref > https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
func HashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	return string(hash), err
}

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
