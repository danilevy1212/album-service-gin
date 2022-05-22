package main

import (
	"github.com/danilevy1212/album-service-gin/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	albums_routes := r.Group("/albums")
	{
		albums_routes.GET("", controllers.GetAll)
		albums_routes.POST("", controllers.Create)
		albums_routes.GET("/:id", controllers.GetByID)
		albums_routes.DELETE("/:id", controllers.Delete)
		albums_routes.PATCH("/:id", controllers.Patch)
	}

	r.Run(":3000")
}
