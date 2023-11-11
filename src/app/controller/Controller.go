package controller

import "github.com/gin-gonic/gin"

func GetRequestHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"Success": "Ok",
	})
}
