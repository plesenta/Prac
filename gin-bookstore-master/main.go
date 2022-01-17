package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/controllers"
)

func main() {
	r := gin.Default()
	r.POST("/labirint", controllers.Movement)
	r.Run()

}
