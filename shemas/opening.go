package shemas

import (
	"gorm.io/gorm"
	"time"
)

type Opening struct {
	gorm.Model `swaggerignore:"true"`
	Role       string
	Company    string
	Location   string
	Link       string
	Remote     bool
	Salary     int64
}

type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Link      string    `json:"link"`
	Remote    bool      `json:"remote"`
	Salary    int64     `json:"salary"`
}
