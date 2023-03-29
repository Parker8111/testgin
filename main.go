package main

import (
	"tesgin/database"
	"tesgin/han"

	"github.com/gin-gonic/gin"
)

func init() {
	han.Sql = database.Opendata()
}

func main() {
	// han.Val()
	r := gin.Default()
	//註冊中間件
	v1 := r.Group("/v1")
	v1.POST("/v2", han.Tes)
	r.Run(":8080")
}
