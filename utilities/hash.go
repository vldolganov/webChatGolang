package utilities

import (
	"github.com/alexedwards/argon2id"
)

func HashPassword(password string) string {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return err.Error()
	}
	return hash
}

func CheckPasswordHash(password, hash string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false
	}
	return match
}
