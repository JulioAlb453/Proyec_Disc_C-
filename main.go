package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//cambio ramdom

	r.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	r.Run(":8000")
}