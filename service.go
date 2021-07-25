package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Default会返回一个engine对象，该对象已经绑定了比如logger这样的中间件
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK,"Hello gin")
	})
	r.Run(":8000")
}
