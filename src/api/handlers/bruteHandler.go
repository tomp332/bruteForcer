package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tomp332/gobrute/src/crud"
	"github.com/tomp332/gobrute/src/models"
	"github.com/tomp332/gobrute/src/utils"
)

// AddCreds godoc
// @Summary Start a new brute force operation
// @Description Start brute force operation
// @Tags Brute Force
// @Accept json
// @Param credentials body []models.IBruteForceCreate true "IBruteForceCreate"
// @Success 200 {array} models.IBruteForceCreate
// @Failure 400 {object} models.ServerError
// @Failure 500 {object} models.ServerError
// @Router /brute [post]
func StartBruteForce(c echo.Context) error {
	var bruteForceTask []models.IBruteForceCreate
	err := c.Bind(&bruteForceTask)
	if err != nil || bruteForceTask == nil {
		log.Printf("Error binding creds struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for BruteForceDTO schema", err))
	}
	if err != nil {
		log.Printf("Error validating creds struct")
		return c.JSONBlob(http.StatusBadRequest,
			utils.BadRequestError("Validation error for BruteForceDTO schema", err))
	}
	addedTasks, err := crud.BruteForceCrud.Add(bruteForceTask)
	if err != nil {
		return c.JSONBlob(http.StatusInternalServerError,
			utils.BadRequestError("Error adding new brute force task to database", err))
	}
	return c.JSON(http.StatusOK, addedTasks)
}
