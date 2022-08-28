package main

import (
	"apiUser/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	port := ":3000"

	controllerGender := controller.NewControllerGender()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/gender", controllerGender.Select)
	r.POST("/gender", controllerGender.Create)
	r.PATCH("/gender", controllerGender.Update)
	r.DELETE("/gender", controllerGender.Delete)

	log.Printf("listen and serve on 0.0.0.0%v", port)
	r.Run(port)
}
