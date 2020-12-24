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
	// タイトル
	// titles := doc.Find("h2.entry-title") 
	// titles.Each(func(i int, s *goquery.Selection){ 
	// 	fmt.Println(s.Text()) 
	// }) 

	thumbnails := doc.Find("a.entry-thumbnail")
	// 詳細ページのURL
	thumbnails.Each(func(i int, s *goquery.Selection) {
		detailUrl, exist := s.Attr("href")
		if !exist {
			log.Print("not exist href")
		}
		fmt.Println(detailUrl)
	})

	// 画像を一件取得
	images, exist := thumbnails.Find("img").Attr("src")
	if !exist {
		log.Print("not exist src")
	}
	fmt.Println(images)

}