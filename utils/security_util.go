package utils

import "golang.org/x/crypto/bcrypt"

/**
 * := Coded with love by Sakib Sami on 17/11/17.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func MakePassword(password string) string {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(encrypted)
}

func IsPasswordMatched(password string, encrypted string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encrypted), []byte(password))
	return err == nil
}
