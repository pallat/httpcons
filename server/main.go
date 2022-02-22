package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		select {
		case <-ctx.Done():
			fmt.Println("success")
		case <-c.Request.Context().Done():
			fmt.Println("connection lost")
		}
	})

	r.Run()
}
