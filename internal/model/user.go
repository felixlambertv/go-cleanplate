package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	Name      string         `json:"name" gorm:"not null" example:"user name"`
	Email     string         `json:"email" gorm:"not null;unique" example:"email@email.com"`
	Password  string         `json:"-" gorm:"not null" example:"password123"`
	UserLevel uint           `json:"user_level" gorm:"not null" example:"1"`
	CreatedAt time.Time      `json:"createdAt,omitempty" example:"2023-01-01T15:01:00+00:00"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty" example:"2023-02-11T15:01:00+00:00"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
