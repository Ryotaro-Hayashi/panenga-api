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
}

func GetPanels() (panels []Panel, err error) {
	rows, err := db.Conn.Query("SELECT id, title, panel_image, created_at FROM panels")

	if err != nil {
		err = fmt.Errorf("query get panels, %s", err)
		fmt.Println(err)
		return
	}

	for rows.Next() {
		var panel Panel
		err = rows.Scan(&panel.ID, &panel.Title, &panel.PanelImage, &panel.CreatedAt)
		if err != nil {
			err = fmt.Errorf("cast get panels, %s", err)
			fmt.Println(err)
			return
		}
		panels = append(panels, panel)
	}

	return
}
