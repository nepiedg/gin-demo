package main

import (
	"gin-demo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.ApiRoutersInit(r)
	r.Run()
}
