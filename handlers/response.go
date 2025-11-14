package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendErrorResponse(c *gin.Context, code int, message string) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func sendSuccessResponse(c *gin.Context, operation string, data interface{}) {
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Opertation from handler: %s Sucessfull", operation),
		"data":    data,
	})
}
