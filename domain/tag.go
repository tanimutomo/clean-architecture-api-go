package domain

import (
	"time"
)

type Tags []Tag

type Tag struct {
	ID        int        `json:"id" gorm:"column:id;primary_key"`
	Name      string     `json:"name" binding:"required" gorm:"column:name"`
	ArticleID int        `json:"articleid" gorm:"column:article_id"`
	UserID    int        `json:"userid" gorm:"column:user_id"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
	DeletedAt *time.Time `json:"-"`
}
