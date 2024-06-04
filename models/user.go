package models

import (
	"time"

	"gorm.io/gorm"
)

type PreparationType string

const (
    BCS PreparationType = "BCS"
    IBA PreparationType = "IBA"
)

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int `json:"age"`
	Email string `json:"email" gorm:"unique"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`



	Level int `json:"level"`
	DayStreak int `json:"day_streak"`
	Badges []Badge `json:"badges" gorm:"many2many:user_badge;"`

	PreparationType PreparationType `json:"preparation_type"`

	Friends []User `json:"friends" gorm:"many2many:user_friend;"`

}

type Badge struct {
	ID uint `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
	Description string `json:"description"`
	Icon string `json:"icon"`
	Level int `json:"level"`
	Users []User `json:"users" gorm:"many2many:user_badge;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


