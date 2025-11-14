package handlers

import (
	"api/golang/gin-gonic/shemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary      Delete opening
// @Description  Deleta uma vaga existente pelo ID
// @Tags         Openings
// @Param        id   path      int  true  "ID da vaga"
// @Produce      json
// @Success      200  {object}  shemas.Opening
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /opening/{id} [delete]
func DeleteOpeningHandler(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
		return
	}

	var opening shemas.Opening
	if err := db.First(&opening, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	sendSuccessResponse(c, "Deleted-Opening", opening)
}
