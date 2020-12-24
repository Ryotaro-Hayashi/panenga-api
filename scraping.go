package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
)

func main() {
	doc, err := goquery.NewDocument("https://kaopane.com/faces/") 
	if err != nil {
		log.Print("htmlの取得に失敗しました") 
	} 
	タイトル
	titles := doc.Find("h2.entry-title") 
	titles.Each(func(i int, s *goquery.Selection){ 
		fmt.Print(s.Text()) 
	}) 

	// 画像URL
// 	images := doc.Find("a.entry-thumbnail")

// 	func GetPage(url string) {
// 		doc, _ := goquery.NewDocument(url)
// 		doc.Find("img").Each(func(_ int, s *goquery.Selection) {
// 			 url, _ := s.Attr("src")
// 			 fmt.Println(url)
// 		})
//    }

}