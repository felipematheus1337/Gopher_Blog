package handler

import (
	"net/http"
	"strings"

	"github.com/felipematheus1337/GoPHER_Blog/dto"
	"github.com/felipematheus1337/GoPHER_Blog/mapper"
	"github.com/felipematheus1337/GoPHER_Blog/service"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service}
}

func (p *PostHandler) CreatePost(ctx *gin.Context) {

	var postDTO dto.CreatePostDTO

	if err := ctx.ShouldBindJSON(&postDTO); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	errValidate := postDTO.Validate()

	if errValidate != nil {
		msgErro := errValidate.Error()
		errStrings := strings.Split(msgErro, ",")
		ErrParamIsRequired(errStrings[0], errStrings[1])
	}

	post := mapper.ToSchema(postDTO)

	response, err := p.service.CreatePost(post)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, response)

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
