package models

import (
	"time"
)

type SoftSkill struct {
	ID          uint      `gorm:"primaryKey" json:"-"`
	UUID        string    `json:"id"`
	Title       string    `json:"title"`
	Icon        string    `json:"icon"`
	Description string    `json:"description"`
	Order       int       `json:"order"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (SoftSkill) TableName() string {
	return "soft_skills"
}