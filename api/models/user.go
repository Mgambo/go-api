package models

import (
	"database/sql"

	models_utils "github.com/mgambo/go-api/api/models/utils"
	internal_models "github.com/mgambo/go-api/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	Username    string `gorm:"not null"`
	Password    []byte `gorm:"not null"`
	FirstName   string `gorm:"not null"`
	LastName    sql.NullString
	DateOfBirth string `gorm:"not null"`
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
