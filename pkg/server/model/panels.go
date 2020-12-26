package model

import (
	"fmt"
	"panenga-api/pkg/server/db"
)

type Panel struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	PanelImage string `json:"panel_image"`
	CreatedAt  string `json:"create_at"`
	UpdatedAt  string `json:"update_at"`
}

func GetPanels() (panels []Panel, err error) {
	rows, err := db.Conn.Query("SELECT id, title, panel_image, created_at, updated_at FROM panels")

	if err != nil {
		fmt.Errorf("query get panels, %s", err)
		return
	}

	for rows.Next() {
		var panel Panel
		err = rows.Scan(&panel.ID, &panel.Title, &panel.PanelImage, &panel.CreatedAt, &panel.UpdatedAt)
		if err != nil {
			fmt.Errorf("cast get panels, %s", err)
			return
		}
		panels = append(panels, panel)
	}

	return
}
