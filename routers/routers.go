package routers

import (
    "github.com/gin-gonic/gin"

    "gin-starter/controllers/api/v1"
    "gin-starter/pkg/setting"
)

func InitRouter() *gin.Engine {
    r := gin.Default()

    gin.SetMode(setting.Cfg.GetString("run_mode"))

    apiv1 := r.Group("/api/v1")
    {
        //获取标签列表
        apiv1.GET("/contents", v1.GetContent)
    }

    return r
}