package router

import (
	"github.com/felipematheus1337/GoPHER_Blog/handler"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/swag/example/override/docs"
)

func InitializeRoutes(router *gin.Engine, p *handler.PostHandler) {

	basePath := "/api/v1/posts"

	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	RegisterPostRoutes(v1, p)

}

func RegisterPostRoutes(v1 *gin.RouterGroup, p *handler.PostHandler) {

	{
		v1.POST("/", p.CreatePost)                    // POST /api/v1/posts
		v1.PUT("/", p.EditPost)                       // PUT /api/v1/posts
		v1.POST("/publishOrUnpublish", p.PublishPost) // POST /api/v1/posts/publishOrUnpublish
		v1.GET("/", p.ListPosts)                      // GET /api/v1/posts
		v1.GET("/show", p.GetPostById)                // GET /api/v1/posts/show
		v1.DELETE("/", p.DeletePost)

	}
}
