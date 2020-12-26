package server

import (
	"app_data/pkg/server/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// サーバーを起動させるためのエンジンを初期化
var Router *gin.Engine

// 各種パッケージの init関数は main関数よりも先に呼ばれる
func init() {
	// デフォルトのエンジンを作成
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		// 許可するメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
		},
		// 許可するリクエストヘッダー
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可するアクセス元
		AllowOrigins: []string{
			"*",
		},
		MaxAge: 24 * time.Hour,
	}))

	router.GET("/panels", handler.HandleGetPanels)

	Router = router
}