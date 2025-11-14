package handlers

import (
	"api/golang/gin-gonic/shemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary      Update opening
// @Description  Atualiza os dados de uma vaga existente pelo ID
// @Tags         Opening
// @Accept       json
// @Produce      json
// @Param        id       path      int                  true  "ID da vaga"
// @Param        request  body      UpdateOpeningRequest  true  "Dados atualizados da vaga"
// @Success      200      {object}  shemas.Opening
// @Failure      400      {object}  map[string]string
// @Failure      404      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /opening/{id} [put]
func UpdateOpeningHandler(c *gin.Context) {
	var request UpdateOpeningRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	var opening shemas.Opening
	if err := db.First(&opening, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Opening not found"})
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": opening})
}
