package service

import "github.com/gin-gonic/gin"

func CreateExaCustomer(c *gin.Context) {
	c.JSON(200, gin.H{
		"Blog":   "www.flysnow.org",
		"wechat": "flysnow_org",
	})
}
