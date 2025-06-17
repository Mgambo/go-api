package models

import (
	models_utils "github.com/mgambo/go-api/api/models/utils"
	internal_models "github.com/mgambo/go-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username    string  `json:"username"`
	Password    []byte  `json:"password"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	DateOfBirth string  `json:"date_of_birth"`
	*internal_models.DatabaseBaseModel
}

func setPassword(u User, tx gorm.DB) {
	password := []byte(u.Password)
	if pw, err := bcrypt.GenerateFromPassword(password, 10); err == nil {
		tx.Statement.SetColumn("Password", pw)
	}
	return
}

func isMatchPassword(u User, password string) bool {
	err := bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	return err == nil
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	setPassword(*u, *tx)
	models_utils.ToLowercase("Username", u.Username, tx)
	return
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	setPassword(*u, *tx)
	models_utils.ToLowercase("Username", u.Username, tx)
	return
}
