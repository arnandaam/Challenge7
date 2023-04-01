package route

import (
	"book/handler"
	"book/service"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, app service.ServiceInterface) {
	server := handler.NewHttpServer(app)
	api := r.Group("/book")
	{
		api.POST("", server.CreateBook)
		api.GET("/:id", server.GetBookID)
		api.GET("", server.GetBookAll)
		api.DELETE("/:id", server.DeleteBook)
		api.PUT("/:id", server.UpdateBook)
	}
}
