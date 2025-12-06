package schemas

import (
	"time"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string
	Body      string
	Published *bool
	Author    string
	Tags      []string
}

type PostResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Published bool           `json:"published"`
	Author    string         `json:"author"`
	Tags      []string       `json:"tags"`
}
