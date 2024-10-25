package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	types "livecode/internal/pkg/basic_http"

	"github.com/gin-gonic/gin"
)

func main() {
	serverPort, ok := os.LookupEnv("BASIC_SERVER_PORT")
	if !ok {
		fmt.Println("not port provided")
		os.Exit(1)
	}

	engine := gin.New()

	engine.GET("/health/", func(ctx *gin.Context) {
		ctx.Status(http.StatusNoContent)
	})
	engine.POST("/x2", func(ctx *gin.Context) {
		var req types.X2Request
		if err := ctx.BindJSON(&req); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		res := types.X2Response{
			Val: req.Val * 2,
		}

		ctx.JSON(http.StatusOK, res)
	})

	if err := engine.Run(":" + serverPort); err != nil {
		log.Fatal(err)
	}
}
