package router

import (
	"github.com/felipematheus1337/GoPHER_Blog/handler"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/swag/example/override/docs"
)

func initializeRoutes(router *gin.Engine, p *handler.PostHandler) {

	basePath := "/api/v1/posts"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	RegisterPostRoutes(v1, p)

}

func RegisterPostRoutes(v1 *gin.RouterGroup, p *handler.PostHandler) {

	{
		v1.POST(v1.BasePath(), p.CreatePost)

		v1.PUT(v1.BasePath(), p.EditPost)

		v1.POST(v1.BasePath()+"/publish", p.PublishPost)

		v1.PATCH(v1.BasePath()+"/unpublish", p.UnpublishPost)

		v1.GET(v1.BasePath(), p.ListPosts)

		v1.GET(v1.BasePath()+"/show", p.GetPostById)

		v1.DELETE(v1.BasePath(), p.DeletePost)

	}
}
