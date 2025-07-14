package utils

import (
	"errors"
	"strings"
)

func ValiateEmail(email string) error {
	if strings.Count(email, ".") < 1 || strings.Count(email, "@") != 1 {
		return errors.New("email invalido")
	}
	return nil
}
