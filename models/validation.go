package models

import (
	"regexp"
	"unicode"
)

const (
	StdIDLength    = 22
	TokenIDLength  = 66
	MaxEmailLength = 200
	MinEmailLength = 5
)

var (
	usernameRegexp = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,15}\s?[a-zA-Z0-9_-]{1,15}$`)
	emailRegexp    = regexp.MustCompile(`^[a-zA-z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	nameRegexp     = regexp.MustCompile(`^[\sa-zA-Z0-9_-]+$`)
)

func isValidPassword(password string) bool {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
		case unicode.IsUpper(ch):
			uppercasePresent = true
		case unicode.IsLower(ch):
			lowercasePresent = true
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
		case ch == ' ':
			return false
		}
	}

	return numberPresent && uppercasePresent && lowercasePresent && specialCharPresent && len(password) >= 8 && len(password) <= 15
}

func IsValidID(id string) bool {
	return len(id) == StdIDLength
}
