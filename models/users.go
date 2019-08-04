package models

import (
	"github.com/jinzhu/gorm"
	c "gibhub.com/skantuz/curso/controllers"
)


type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique; not null"`
	Password string `json:"password" gorm:"not null"`
	Email string `json:"email"`

}

func (u *User) create() map[string]interface{}{
	GetDB().create(u)
	response := c.Menssage(true,"Usuario Creado")
	response["user"] = u
	return u

}