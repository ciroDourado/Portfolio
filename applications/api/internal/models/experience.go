package models

import "time"

type Experience struct {
	ID          uint       `gorm:"primaryKey" json:"-"`
	UUID        string     `json:"id"`
	CompanyName string     `json:"company_name"`
	Role        string     `json:"role"`
	Description *string    `json:"description"`
	Results     []string   `gorm:"serializer:json" json:"results"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	IsCurrent   bool       `json:"is_current"`
}

func (Experience) TableName() string {
	return "experiences"
}