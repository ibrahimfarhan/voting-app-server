package models

import (
	"strings"
	"time"

	"github.com/ibrahimfarhan/voting-app/voting-app-server/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         string    `json:"id" bson:"_id"`
	OAuthID    string    `json:"-" bson:"oauthId"`
	Username   string    `json:"username" bson:"username"`
	Name       string    `json:"name" bson:"name"`
	Email      string    `json:"-" bson:"email"`
	Password   string    `json:"-" bson:"password"`
	PictureURL string    `json:"pictureUrl" bson:"pictureUrl"`
	Teams      []*Team   `json:"teams,omitempty" bson:"-"`
	CreatedAt  time.Time `json:"-" bson:"createdAt"`
	UpdatedAt  time.Time `json:"-" bson:"updatedAt"`
}

func (u *User) Presave(r *RegisterData) error {
	u.Email = strings.ToLower(r.Email)
	u.Username = r.Username
	u.ID = utils.NewID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	var err error
	u.Password, err = HashPassword(r.Password)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func IsCorrectPassword(hash, password string) bool {
	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
