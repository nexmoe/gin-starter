package v1

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)

func GetContent(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
        "code" : 200,
        "msg" : "和",
        "data" : "2333",
    })
}