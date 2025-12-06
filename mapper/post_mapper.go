package mapper

import (
	"github.com/felipematheus1337/GoPHER_Blog/dto"
	"github.com/felipematheus1337/GoPHER_Blog/schemas"
)

func CreateToSchema(dto dto.CreatePostDTO) *schemas.Post {
	return &schemas.Post{
		Title:     dto.Title,
		Body:      dto.Body,
		Published: dto.Published,
		Author:    dto.Author,
		Tags:      dto.Tags,
	}
}

func UpdateToSchema(dto dto.UpdatePostDTO) *schemas.Post {
	return &schemas.Post{
		Title:  dto.Title,
		Body:   dto.Body,
		Author: dto.Author,
		Tags:   dto.Tags,
	}
}
