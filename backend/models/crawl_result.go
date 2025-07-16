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
    H4Count           int            `json:"h4_count"`
	H5Count           int            `json:"h5_count"`
	H6Count           int            `json:"h6_count"`
	InternalLinks     int            `json:"internal_links"`
	ExternalLinks     int            `json:"external_links"`
	InaccessibleLinks int            `json:"inaccessible_links"`
	HasLoginForm      bool           `json:"has_login_form"`
    Status            string         `gorm:"default:'queued'"`
	UserID            uint           `json:"user_id"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`
}

/* 
	URL               string
	HTMLVersion       string
	Title             string
	H1Count           int
	H2Count           int
	H3Count           int
	H4Count           int
	H5Count           int
	H6Count           int
	InternalLinks     int
	ExternalLinks     int
	InaccessibleLinks int
	HasLoginForm      bool
	Status            string `gorm:"default:'queued'"`
	UserID            uint */