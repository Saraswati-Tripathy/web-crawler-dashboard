package models

import (
	"gorm.io/gorm"
	"time"
)

type CrawlResult struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	URL               string         `json:"url"`
	Title             string         `json:"title"`
	HTMLVersion       string         `json:"html_version"`
	H1Count           int            `json:"h1_count"`
	H2Count           int            `json:"h2_count"`
	H3Count           int            `json:"h3_count"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	HasLoginForm      bool           `json:"has_login_form"`
	UserID            uint           `json:"user_id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}
