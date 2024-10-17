package main

import (
	"log"
	"net/http"
	"time"

	types "livecode/internal/pkg/basic_http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	engine.GET("/health/", func(ctx *gin.Context) {
		time.Sleep(30 * time.Second)
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

	go func() {
		if err := engine.Run(":7500"); err != nil {
			log.Fatal(err)
		}
	}()
}
