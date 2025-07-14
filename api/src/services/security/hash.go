package security

import "golang.org/x/crypto/bcrypt"

func Hash(str string) ([]byte, error) {
	b := []byte(str)
	return bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
}

func Check(hashed, str string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(str))
}
