package v1

import (
	"ginProject/service"
	"github.com/gin-gonic/gin"
)

func CreateExaCustomer(c *gin.Context) {
	service.CreateExaCustomer(c)
}
