package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" validate="required,min=2,max=50"`
	Email    string `json:"email" validate="email,required"`
	Username string `gorm:"unique" json:"username" validate="required,min=2,max=50"`
	Password string `json:"password" validate="required,min=6,max=50"`
	// UserType     string `json:"role" validate="required, eq=ADMIN|eq=USER"`
}

type Credentials struct {
	Id       string `json: "id" validate="required, min=2, max=50"`
	Password string `json: "password" validate="required, min=6, max=50"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) CreateUser() *User {
	db.NewRecord(u)
	db.Create(&u)
	return u
}

// TODO: Melhorar funções
func GetUserByUsername(username string) (*User, error) {
	var u User
	if err := db.Where("username = ?", username).First(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserById(id string) (*User, error) {
	var u User

	results := db.Where("username = ? OR email = ?", id, id).First(&u)
	if results.Error != nil || results.RowsAffected < 1 {
		return nil, fmt.Errorf("invalid username/email")
	}
	return &u, nil
}
