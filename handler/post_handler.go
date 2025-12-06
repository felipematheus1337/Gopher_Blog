package handler

import (
	"net/http"
	"strconv"
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

	post := mapper.CreateToSchema(postDTO)

	response, err := p.service.CreatePost(post)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(ctx, "create-post", response, http.StatusCreated)

}

func (p *PostHandler) EditPost(ctx *gin.Context) {

	request := dto.UpdatePostDTO{}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, isValid := GetIdFromQuery(ctx)

	if !isValid {
		sendError(ctx, http.StatusBadRequest, "Id inválido.")
		return
	}

	response, err := p.service.UpdatePost(id, &request)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSucess(ctx, "update-post", response, http.StatusOK)

}

func (p *PostHandler) PublishPost(ctx *gin.Context) {
	id, isValid := GetIdFromQuery(ctx)

	if !isValid {
		sendError(ctx, http.StatusBadRequest, "ID inválido")
		return
	}

	strPublished := ctx.Query("published")
	isPublished, err := strconv.ParseBool(strPublished)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "Parâmetro 'published' inválido")
		return
	}

	err = p.service.PublishPost(id, isPublished)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "Post atualizado com sucesso", id, http.StatusNoContent)
}

func (p *PostHandler) ListPosts(context *gin.Context) {

}

func (p *PostHandler) GetPostById(ctx *gin.Context) {
	id, isValid := GetIdFromQuery(ctx)

	if !isValid {
		sendError(ctx, http.StatusBadRequest, "ID inválido")
		return
	}

	response, err := p.service.FindById(id)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error finding post by id.")
		return
	}

	sendSucess(ctx, "get-post", response, http.StatusOK)
}

func (p *PostHandler) DeletePost(ctx *gin.Context) {

}
