package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	app := gin.Default()

	app.GET("", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"message": os.Getenv("MESSAGE"),
		})
	})

	app.GET("/echo/{message}", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"message": ctx.Param("message"),
		})
	})

	app.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
