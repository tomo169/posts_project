package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Title     string         `gorm:"type:varchar(100)" json:"title"`
	Content   string         `gorm:"type:text" json:"content"`
	AuthorID  uint           `json:"authorId"`
	Author    User           `gorm:"foreignKey:AuthorID" json:"-"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
