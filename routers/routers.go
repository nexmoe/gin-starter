package routers

import (
	"github.com/gin-gonic/gin"

	v1 "gin-starter/controllers/api/v1"
	"gin-starter/pkg/setting"
)

func InitRouter() *gin.Engine {
	gin.SetMode(setting.Cfg.GetString("run_mode"))

	r := gin.Default()

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/contents", v1.GetContent)
	}

	return r
}
