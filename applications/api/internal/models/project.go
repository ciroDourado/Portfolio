package models

import (
	"time"
)

func (Project) TableName() string {
	return "projects"
}

type Project struct {
	ID               uint       `gorm:"primaryKey" json:"-"`
	UUID             string     `json:"id"`
	Title            string     `json:"title"`
	Slug             string     `json:"slug"`
	ShortDescription *string    `json:"short_description"`
	Content          *string    `json:"content"`
	RepositoryURL    *string    `json:"repository_url"`
	LiveURL          *string    `json:"live_url"`
	ThumbnailURL     *string    `json:"thumbnail_url"`
	StartedAt        *time.Time `json:"started_at"`
	FinishedAt       *time.Time `json:"finished_at"`
	IsFeatured       bool       `json:"is_featured"`
	Skills           []Skill    `gorm:"many2many:project_skill" json:"skills"`
}