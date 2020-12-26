package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"app_data/pkg/server/db"
	"strings"
)

func main() {
	doc, err := goquery.NewDocument("https://kaopane.com/faces/")
	if err != nil {
		fmt.Println(fmt.Errorf("go scraping, %s", err))
	}

	// 記事
	articles := doc.Find("article")
	articles.Each(func(_ int, article *goquery.Selection) {
		// タイトル
		title := article.Find("h2.entry-title")

		// サムネイル
		thumbnail := article.Find("a.entry-thumbnail")

		// パネル画像
		panelImage, exist := thumbnail.Find("img").Attr("src")
		if !exist {
			fmt.Println("not exist src")
			return
		}

		slice := strings.Split(title.Text(), "\n\t\t\t")
		slice = strings.Split(slice[1], "\n\t\t")

		// 画像とタイトルを挿入
		_, err := db.Conn.Query("INSERT INTO panels (title, panel_image) VALUES (?, ?)", slice[0], panelImage)

		if err != nil {
			fmt.Println(fmt.Errorf("query post panels, %s, %s", slice[0], err))
		}

		return
	})
}
