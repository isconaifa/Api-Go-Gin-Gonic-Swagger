package handlers

import (
	"api/golang/gin-gonic/shemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @BasePath /api/v1

// @Summary      Create opening
// @Description  Cria uma nova vaga de emprego
// @Tags         Opening
// @Accept       json
// @Produce      json
// @Param        request  body      CreateOpeningRequest  true  "Dados da nova vaga"
// @Success      201      {object}  shemas.Opening
// @Failure      400      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /opening [post]
func CreateOpeningHandler(c *gin.Context) {
	request := CreateOpeningRequest{}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	if err := request.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	opening := shemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Link:     request.Link,
		Salary:   request.Salary,
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if err := db.Create(&opening).Error; err != nil {
		sendErrorResponse(c, http.StatusInternalServerError, err.Error())
		logger.Errorf("Error to create opening: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	sendSuccessResponse(c, "Created-Opening", opening)
}
