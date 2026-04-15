package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ContactMessage struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	UUID      string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Message   string    `json:"message"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ContactMessage) TableName() string {
	return "contact_messages"
}

func (cm *ContactMessage) BeforeCreate(tx *gorm.DB) error {
	if cm.UUID == "" {
		cm.UUID = uuid.New().String()
	}

	return nil
}