package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.Default()
	router.Use(gin.Logger())
	initializeRoutes(router)
	_ = router.Run(":8080")
}
