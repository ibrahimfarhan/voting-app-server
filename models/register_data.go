package models

import (
	"encoding/json"
)

type RegisterData struct {
	LoginData
}

func ValidateRegisterData(registerData []byte) (bool, *RegisterData) {
	var d RegisterData

	err := json.Unmarshal(registerData, &d)
	if err != nil {
		return false, nil
	}

	return usernameRegexp.MatchString(d.Username) && emailRegexp.MatchString(d.Email) && len(d.Email) <= MaxEmailLength && len(d.Email) >= MinEmailLength && isValidPassword(d.Password), &d
}
