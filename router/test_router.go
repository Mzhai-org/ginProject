package router

import (
	"ginProject/app/v1"
	"github.com/gin-gonic/gin"
)

func InitCustomerRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("customer")
	{
		ApiRouter.GET("create", v1.CreateExaCustomer) // 创建客户
	}
}
