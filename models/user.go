package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string `gorm:"firstname"`
	Last_name  string `gorm:"lastname"`
	Email      string `gorm:"unique"`
	Password   string `gorm:"password"`
	Age        int    `gorm:"age"`
	Phone_no   int    `gorm:"phone_no"`
	Role_id    int    `gorm:"role_id"`
}