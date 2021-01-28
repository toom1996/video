package main

import "github.com/gin-gonic/gin"

//test1
func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"data": "success",
		})
	})
	_ = r.Run()
}
