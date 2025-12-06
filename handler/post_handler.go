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

	sendSuccess(ctx, "create-post", response, http.StatusCreated)

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

	sendSuccess(ctx, "update-post", response, http.StatusOK)

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

func (p *PostHandler) ListPosts(ctx *gin.Context) {

	published := ctx.DefaultQuery("published", "")
	author := ctx.DefaultQuery("author", "")
	tag := ctx.DefaultQuery("tag", "")
	search := ctx.DefaultQuery("search", "")
	page := ctx.DefaultQuery("page", "1") // Página padrão é 1
	limit := ctx.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt <= 0 {
		limitInt = 10
	}

	var publishedBool *bool
	if published != "" {
		b, err := strconv.ParseBool(published)
		if err != nil {
			sendError(ctx, http.StatusBadRequest, "Parâmetro 'published' inválido")
			return
		}
		publishedBool = &b
	}

	posts, total, err := p.service.ListPosts(publishedBool, author, tag, search, pageInt, limitInt)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "list-posts",
		gin.H{"data": posts,
			"total":       total,
			"page":        pageInt,
			"limit":       limitInt,
			"total_pages": (total + int64(limitInt) - 1) / int64(limitInt)},
		http.StatusOK)

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

	sendSuccess(ctx, "get-post", response, http.StatusOK)
}

func (p *PostHandler) DeletePost(ctx *gin.Context) {
	id, isValid := GetIdFromQuery(ctx)

	if !isValid {
		sendError(ctx, http.StatusBadRequest, "ID inválido")
		return
	}

	err := p.service.DeletePost(id)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "Error deleting post.")
		return
	}

	sendSuccess(ctx, "Post deletado com sucesso.", id, http.StatusOK)

}
