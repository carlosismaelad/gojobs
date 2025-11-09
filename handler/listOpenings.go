package handler

import (
	"net/http"

	"github.com/carlosismaelad/gojobs/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary List Openings
// @Description List all job openings
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failures 400 {object} ErrorResponse
// @Failures 404 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}