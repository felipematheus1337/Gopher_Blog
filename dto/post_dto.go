package dto

import "github.com/lib/pq"

type CreatePostDTO struct {
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Published *bool          `json:"published,omitempty"`
	Author    string         `json:"author"`
	Tags      pq.StringArray `json:"tags"`
}
