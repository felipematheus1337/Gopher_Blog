package dto

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
)

type CreatePostDTO struct {
	Title     string         `json:"title"`
	Body      string         `json:"body"`
	Published *bool          `json:"published,omitempty"`
	Author    string         `json:"author"`
	Tags      pq.StringArray `json:"tags"`
}

type UpdatePostDTO struct {
	Title  string         `json:"title"`
	Body   string         `json:"body"`
	Tags   pq.StringArray `json:"tags"`
	Author string         `json:"author"`
}

func (d *CreatePostDTO) Validate() error {
	// Validações
	if d.Title == "" {
		return errors.New(fmt.Sprintf(" 'title', 'string'"))
	}

	if d.Body == "" {
		return errors.New(fmt.Sprintf("'body','string'"))
	}

	if d.Published != nil && *d.Published {
		return errors.New(fmt.Sprintf("'published','bool'"))
	}

	if d.Author == "" {
		return errors.New(fmt.Sprintf("'author','string'"))
	}

	if d.Tags == nil {
		d.Tags = pq.StringArray{}
	}

	if d.Published == nil {
		var falso = false
		d.Published = &falso
	}

	return nil
}

func (u *UpdatePostDTO) Validate() error {
	if u.Title == "" {
		return errors.New(fmt.Sprintf(" 'title', 'string'"))
	}
	if u.Body == "" {
		return errors.New(fmt.Sprintf(" 'body', 'string'"))
	}
	if u.Tags == nil {
		u.Tags = pq.StringArray{}
	}
	if u.Author == "" {
		return errors.New(fmt.Sprintf(" 'author', 'string'"))
	}
	return nil
}
