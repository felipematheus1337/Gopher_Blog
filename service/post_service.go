package service

import (
	"github.com/felipematheus1337/GoPHER_Blog/schemas"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func (s PostService) CreatePost(post *schemas.Post) (*schemas.PostResponse, error) {

	var response *schemas.PostResponse

	if err := s.db.Create(&post).Error; err != nil {
		return &schemas.PostResponse{}, err
	}

	response = &schemas.PostResponse{
		ID:        post.ID,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
		DeletedAt: post.DeletedAt,
		Tags:      post.Tags,
		Title:     post.Title,
		Body:      post.Body,
		Published: *post.Published,
		Author:    post.Author,
	}

	return response, nil
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db}
}
