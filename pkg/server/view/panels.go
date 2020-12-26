package view

import (
	"app_data/pkg/server/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PanelsGetResponse struct {
	Panels []model.Panel `json:"panels"`
}

type Error struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func ResponseGetPanels(panels []model.Panel, c *gin.Context) {
	response := &PanelsGetResponse{
		Panels: panels,
	}
	c.JSON(http.StatusOK, response)
}

func ErrorResponse(code int, msg, desc string, c *gin.Context) {
	body := Error{
		Code:        code,
		Message:     msg,
		Description: desc,
	}
	c.JSON(code, body)
}
