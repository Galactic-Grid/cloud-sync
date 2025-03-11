package model

import "time"

type GitConfig struct {
	ID        uint   `gorm:"primarykey"`
	RepoName  string `gorm:"type:varchar(255);not null"`
	RepoURL   string `gorm:"type:varchar(255);not null"`
	Token     string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
