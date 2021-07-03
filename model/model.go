package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"size:15; not null; unique" json:"username"`
	Password string `gorm:"size:20; not null" json:"password"`
	// Email    string `gorm:"size:20; not null; unique" json:"email"`
	// CityList string `gorm:"size:255;" json:"city_list"`
}
