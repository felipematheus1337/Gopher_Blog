package mapper

import (
	"github.com/felipematheus1337/GoPHER_Blog/dto"
	"github.com/felipematheus1337/GoPHER_Blog/schemas"
)

func ToSchema(dto dto.CreatePostDTO) *schemas.Post {
	return &schemas.Post{
		Title:     dto.Title,
		Body:      dto.Body,
		Published: dto.Published,
		Author:    dto.Author,
		Tags:      dto.Tags,
	}
}
