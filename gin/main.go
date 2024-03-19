package main

import (
	"gin/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/mahasiswa", controller.InsertMhs)
	r.GET("/mahasiswa", controller.GetMhs)
	r.PUT("/mahasiswa/:id", controller.UpdateMhs)
	r.DELETE("/mahasiswa/:id", controller.DeleteMhs)
	r.Run(":8888")
}
