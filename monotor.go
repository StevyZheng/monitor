package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	v1_0 := router.Group("/v1.0")
	v1_0.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"app":     "Roycom",
			"version": "v1.0",
		})
	})
	ser := &http.Server{
		Addr:           fmt.Sprintf(":%d", 1991),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = ser.ListenAndServe()
}