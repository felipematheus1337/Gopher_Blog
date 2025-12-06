package service

import (
	"log"
	"strconv"

	"github.com/felipematheus1337/GoPHER_Blog/dto"
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

func (s PostService) UpdatePost(id string, request *dto.UpdatePostDTO) (*schemas.PostResponse, error) {

	idInt, err := strconv.Atoi(id)

	if err != nil {
		return &schemas.PostResponse{}, err
	}

	postSchema := s.getById(uint(idInt))

	postToSave := updateProperties(postSchema, request)

	var postResponse *schemas.PostResponse

	if err := s.db.Save(&postToSave).Error; err != nil {
		postResponse := schemas.PostResponse{
			ID:        postToSave.ID,
			CreatedAt: postToSave.CreatedAt,
			UpdatedAt: postToSave.UpdatedAt,
			DeletedAt: postToSave.DeletedAt,
			Tags:      postToSave.Tags,
			Title:     postToSave.Title,
			Body:      postToSave.Body,
			Published: *postToSave.Published,
			Author:    postToSave.Author,
		}
		return &postResponse, err
	}

	return postResponse, nil
}

func updateProperties(schema *schemas.Post, request *dto.UpdatePostDTO) *schemas.Post {
	return &schemas.Post{
		Title:     request.Title,
		Body:      request.Body,
		Tags:      request.Tags,
		Author:    request.Author,
		Published: schema.Published,
	}
}

func NewPostService(db *gorm.DB) *PostService {
	return &PostService{db}
}

func (s *PostService) getById(id uint) *schemas.Post {
	var post schemas.Post

	if err := s.db.First(&post, id).Error; err != nil {
		log.Fatal("Error while fetching post by id")
		return &schemas.Post{}
	}

	return &post

}
