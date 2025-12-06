package handler

import "github.com/gin-gonic/gin"

type PostHandler struct {
}

func NewPostHandler() *PostHandler {
	return &PostHandler{}
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
