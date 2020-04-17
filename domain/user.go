package domain

import (
	"time"
)

type Users []User

type User struct {
	ID        int        `json:"id" gorm:"column:id;primary_key"`
	Name      string     `json:"name" binding:"required" gorm:"column:name"`
	Password  string     `json:"password" binding:"required" gorm:"column:password"`
	Email     string     `json:"email" binding:"required" gorm:"column:email"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}

type LoginUser struct {
	ID       int    `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
