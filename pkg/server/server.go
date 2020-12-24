package server

import (
	"github.com/gin-gonic/gin"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	r := gin.Default()
	r.Run()
}
