package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashed), nil
}

func CheckPasswordHash(password string, hash_password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash_password), []byte(password))
	return err == nil
}
