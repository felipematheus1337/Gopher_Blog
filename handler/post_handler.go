package handler

import (
	"github.com/felipematheus1337/GoPHER_Blog/service"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service}
}

func (p *PostHandler) CreatePost(c *gin.Context) {

}

func (p *PostHandler) EditPost(context *gin.Context) {

}

func (p *PostHandler) PublishPost(context *gin.Context) {

}

func (p *PostHandler) UnpublishPost(context *gin.Context) {

}

func (p *PostHandler) ListPosts(context *gin.Context) {

}

func (p *PostHandler) GetPostById(context *gin.Context) {

}

func (p *PostHandler) DeletePost(context *gin.Context) {

}
