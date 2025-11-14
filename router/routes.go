package router

import (
	docs "api/golang/gin-gonic/docs"
	"api/golang/gin-gonic/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	handlers.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	var v1 *gin.RouterGroup = router.Group(basePath)
	{
		v1.POST("/opening", handlers.CreateOpeningHandler)
		v1.GET("/openings", handlers.ListOpeningHandler)
		v1.GET("/opening/:id", handlers.ShowOpeningHandler)
		v1.PUT("/opening/:id", handlers.UpdateOpeningHandler)
		v1.DELETE("/opening/:id", handlers.DeleteOpeningHandler)
	}
	//Initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	var v2 *gin.RouterGroup = router.Group("/api/v2")
	{
		v2.GET("/opening", func(c *gin.Context) {})
		v2.POST("/opening", func(c *gin.Context) {})
	}
}
