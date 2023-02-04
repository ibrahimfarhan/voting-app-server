package models

import "encoding/json"

type LoginData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ValidateLoginData(loginData []byte) (bool, *LoginData) {
	var l LoginData
	err := json.Unmarshal(loginData, &l)
	if err != nil {
		return false, nil
	}

	return (usernameRegexp.MatchString(l.Username) || emailRegexp.MatchString(l.Email)) && isValidPassword(l.Password), &l
}
