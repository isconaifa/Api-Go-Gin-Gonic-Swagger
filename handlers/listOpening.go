package handlers

import (
	"api/golang/gin-gonic/shemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary      List openings
// @Description  Retorna todas as vagas de emprego cadastradas
// @Tags         Openings
// @Produce      json
// @Success      200  {array}   shemas.Opening
// @Failure      500  {object}  map[string]string
// @Router       /openings [get]
func ListOpeningHandler(c *gin.Context) {
	var openings []shemas.Opening

	if err := db.Find(&openings).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccessResponse(c, "Opening", openings)
}
