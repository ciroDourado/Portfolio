package models

type Skill struct {
	ID              uint   `gorm:"primaryKey" json:"-"`
	UUID            string `json:"id"`
	SkillCategoryID uint   `json:"-"`
    Category        *Category `gorm:"foreignKey:SkillCategoryID" json:"category,omitempty"`
	Name             string `json:"name"`
	Slug             string `json:"slug"`
	IconURL          string `json:"icon_url"`
	ProficiencyLevel int    `json:"proficiency_level"`
	IsHighlight      bool   `json:"is_highlight"`
}

func (Skill) TableName() string {
	return "skills"
}

type Category struct {
    ID    uint   `gorm:"primaryKey" json:"-"`
    Name  string `json:"name"`
    Slug  string `json:"slug"`
    Order int    `json:"order"`
}

func (Category) TableName() string {
    return "skill_categories"
}