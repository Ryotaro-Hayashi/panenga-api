package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	doc, err := goquery.NewDocument("https://kaopane.com/faces/") 
	if err != nil {
		log.Print(err) 
	} 

	// 記事
	articles := doc.Find("article")
	articles.Each(func(_ int, article *goquery.Selection) {
		// タイトル
		title := article.Find("h2.entry-title")
		fmt.Println(title.Text())

		// サムネイル
		thumbnail := article.Find("a.entry-thumbnail")

		// パネル画像
		panelImage, exist := thumbnail.Find("img").Attr("src")
		if !exist {
			log.Print("not exist src")
		}
		fmt.Println(panelImage)
	})
}