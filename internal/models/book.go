package models

import "time"

type Book struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	Genre         string    `json:"genre"`
	PublishedYear int       `json:"published_year"`
	Available     bool      `json:"available"`
	CreatedAt     time.Time `json:"created_at"`
}
