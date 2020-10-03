package routers

import (
    "github.com/gin-gonic/gin"

    "gin-starter/controllers/api/v1"
    "gin-starter/pkg/setting"
)

func InitRouter() *gin.Engine {
    r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    apiv1 := r.Group("/api/v1")
    {
        //获取标签列表
        apiv1.GET("/tags", v1.GetContent)
    }

    return r
}