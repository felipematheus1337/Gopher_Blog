package service

import (
	"github.com/felipematheus1337/GoPHER_Blog/schemas"
	"gorm.io/gorm"
)

type PostService struct {
	db *gorm.DB
}

func (s PostService) CreatePost(post *schemas.Post) (interface{}, error) {

}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db}
}
