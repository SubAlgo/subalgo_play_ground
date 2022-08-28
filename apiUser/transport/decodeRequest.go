package transport

import "github.com/gin-gonic/gin"

func DecodeRequestJson(c *gin.Context, req interface{}) error {
	return c.BindJSON(&req)
}
