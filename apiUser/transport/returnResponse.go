package transport

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ReturnOK(c *gin.Context, res interface{}) {
	if res == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Success",
		})
	} else {
		c.JSON(http.StatusInternalServerError, res)
	}
}

func ReturnBadRequest(c *gin.Context, res interface{}) {
	if res == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "bad request",
		})
	} else {
		c.JSON(http.StatusInternalServerError, res)
	}
}

func ReturnInternalServerError(c *gin.Context, res interface{}) {
	if res == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
	} else {
		c.JSON(http.StatusInternalServerError, res)
	}
}
