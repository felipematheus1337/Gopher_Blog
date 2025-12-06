package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getIdFromQuery(ctx *gin.Context) (string, bool) {
	id := ctx.Query("id")

	if id == "" {
		sendError(ctx, http.StatusBadRequest, "ID IS required")
		return "", false
	}

	return id, true
}
