package domain

import (
	"time"
)

type Articles []Article

type Article struct {
	ID        int        `json:"id" gorm:"column:id;primary_key"`
	Title     string     `json:"title" binding:"required" gorm:"column:id"`
	Content   string     `json:"content" binding:"required" gorm:"column:content"`
	UserID    int        `json:"userid" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
