package models

import "gorm.io/gorm"

type Contact struct {
	gorm.Model
	Nom   string `json:"nom" gorm:"type:text"`
	Email string `json:"email" gorm:"type:text"`
}
