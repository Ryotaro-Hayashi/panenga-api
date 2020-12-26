package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"panenga-api/pkg/server/model"
	"panenga-api/pkg/server/view"
)

func HandleGetPanels(c *gin.Context) {

	panels, err := model.GetPanels()
	if err != nil {
		view.ErrorResponse(
			http.StatusInternalServerError,
			"Internal Server Error",
			err.Error(),
			c,
		)
		return
	}

	view.ResponseGetPanels(panels, c)
}
