package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := SetupRouter()
	err := r.Run(":80")

	if err != nil {
		panic(err)
	}
}

type ResponseData struct {
	Message string `json:"message"`
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})

	if err != nil {
		return nil
	}

	r.GET("/ping", func(c *gin.Context) {
		pong(c)
	})

	return r
}

func pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, ResponseData{
		Message: "pong",
	})
}
