package handlers

import (
	"api/golang/gin-gonic/shemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary      Get opening by ID
// @Description  Retorna uma vaga de emprego específica pelo ID
// @Tags         Openings
// @Produce      json
// @Param        id   path      int  true  "ID da vaga"
// @Success      200  {object}  shemas.Opening
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /opening/{id} [get]
func ShowOpeningHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
		return
	}

	var opening shemas.Opening
	if err := db.Where("id = ?", id).First(&opening).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Opening not found"})
		return
	}

	sendSuccessResponse(c, "Opening", opening)
}
