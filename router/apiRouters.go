package router

import (
	"gin-demo/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/list", api.ApiController{}.List)
		apiRouters.GET("/create", api.ApiController{}.Create)
	}
}
