package api

import (
	"fmt"
	"gin-demo/models"
	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(c *gin.Context) {
	c.String(200, "我是一个api接口")
}
func (con ApiController) List(c *gin.Context) {
	adminLogList := []models.AdminLog{}
	models.DB.Find(&adminLogList)
	c.JSON(200, gin.H{
		"data": adminLogList,
	})
}
func (con ApiController) Create(c *gin.Context) {
	info := models.AdminLog{
		AdminId:    10,
		Title:      "测试添加",
		Url:        "/api/index",
		Ip:         "192.168.1.1",
		Content:    "测试添加内容",
		Createtime: 0,
	}
	models.DB.Create(&info)
	fmt.Println(info)
	c.JSON(200, gin.H{
		"msg":  "新增成功",
		"code": 1,
	})
}
