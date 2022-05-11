package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello/*world", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, world")
	})

	r.Run(":8989")
}
