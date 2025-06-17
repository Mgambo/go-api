package models_utils

import (
	"gorm.io/gorm"
)

func ToLowercase(columnName string, value string, tx *gorm.DB) {
	tx.Statement.SetColumn(columnName, value)
	return
}
