package handler

import (
	"fmt"
	"net/http"

	"github.com/carlosismaelad/gojobs/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Delete Opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failures 400 {object} ErrorResponse
// @Failures 404 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	
	if id == "" {
		sendError(
			ctx, 
			http.StatusBadRequest, 
			errorParamIsRequired("id", "queryParameter").Error(),
		)
		return
	}

	opening := schemas.Opening{}

	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	
	if err := db.Delete(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}
	sendSuccess(ctx, "delete-opening", opening)
}