package view

import (
	"app_data/pkg/server/model"
	"github.com/gin-gonic/gin"
)

type PanelsGetResponse struct {
	Panels []model.Panel `json:"panels"`
}

func ResponseGetPanels(panels []model.Panel, c *gin.Context) {
	response := &PanelsGetResponse{
		Panels: panels,
	}
	c.JSON(200, response)
}
