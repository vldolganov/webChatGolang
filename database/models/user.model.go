package models

import (
	"time"
)

type Users struct {
	ID        uint      `json:"id" gorm:"autoIncrement" gorm:"foreignKey" gorm:"references:UserId"`
	Login     string    `json:"login" gorm:"unique"`
	Password  string    `json:"password"`
	Email     string    `json:"email" gorm:"unique"`
	Avatar    string    `json:"avatar"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"createdAt"`
}
