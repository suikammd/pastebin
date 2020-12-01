package main

import "time"

type Paste struct {
	ShortLink           string    `json:"short_link" gorm:"column:short_link;primary_key"`
	CreatedAt           time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	ExpirationInMinutes int64     `json:"expiration_in_minutes,omitempty" gorm:"column:expiration_in_minutes"`
	Path                string    `json:"path,omitempty" gorm:"column:path"`
	Count               int64     `json:"count,omitempty" gorm:"column:count"`
}
