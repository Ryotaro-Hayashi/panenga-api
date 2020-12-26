package handler

import (
	"app_data/pkg/server/model"
	"app_data/pkg/server/view"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleGetPanels(c *gin.Context) {

	panels, err := model.GetPanels()
	if err != nil {
		view.ErrorResponse(
			http.StatusInternalServerError,
			"Internal Server Error",
			"Database related error",
			c,
		)
		return
	}

	view.ResponseGetPanels(panels, c)
}
