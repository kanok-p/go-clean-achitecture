package password

import "golang.org/x/crypto/bcrypt"

func Encrypt(password string) (passEncrypted string, err error) {
	pBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(pBytes), nil
}
