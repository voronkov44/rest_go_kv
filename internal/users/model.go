package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"name"`
	Email    string `json:"email" gorm:"email"`
	Age      int    `json:"age" gorm:"age"`
	Password string `json:"password" gorm:"password"`
}

func NewUser(name string, email string, age int) *User {
	return &User{
		Name:  name,
		Email: email,
		Age:   age,
	}
}
