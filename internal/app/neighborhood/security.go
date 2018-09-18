package neighborhood

import (
	"golang.org/x/crypto/bcrypt"
)

func hash(val string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(val), 14)
	return string(bytes), err
}

func checkHash(target string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(target))
	return err == nil
}
