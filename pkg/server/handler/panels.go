package handler

import (
	"app_data/pkg/server/model"
	"app_data/pkg/server/view"
	"github.com/gin-gonic/gin"
)

func HandleGetPanels(c *gin.Context) {

	panels, err := model.GetPanels()
	if err != nil {
		c.Status(500)
		return
	}

	view.ResponseGetPanels(panels, c)
}