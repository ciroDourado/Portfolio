package models

type Profile struct {
	ID          uint    `gorm:"primaryKey" json:"-"`
	UUID        string  `json:"id"`
    Name        string  `json:"name"`
    Headline    string  `json:"headline"`
    Bio         *string `json:"bio"`
    AvatarURL   *string `json:"avatar_url"`
    GithubURL   *string `json:"github_url"`
    LinkedinURL *string `json:"linkedin_url"`
    CVURL       *string `json:"cv_url"`
}
