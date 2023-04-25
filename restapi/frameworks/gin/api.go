package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type A struct {
	Name string
}

type Handler struct {
}

func main() {
	engine := gin.Default()
	routerGroup := engine.Group("")
	h := Handler{}
	routerGroup.GET("/digi/:data", h.PostHandle())
	routerGroup.POST("/digi", h.PostHandle())
	engine.Run()
}

func (h Handler) GetHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println(ctx.Query("id"))
		fmt.Println(ctx.Param("data"))
		ctx.JSON(200, `{"message":"get success"}`)
	}
}

func (h Handler) PostHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var a A
		err := ctx.Bind(&a)
		fmt.Printf("error in req: %v", err)
		aBytes, err := json.Marshal(a)
		fmt.Println(string(aBytes), err)
		ctx.JSON(200, `{"message":"post success"}`)
	}
}
