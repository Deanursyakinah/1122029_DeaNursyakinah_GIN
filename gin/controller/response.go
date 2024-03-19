package controller

import (
	m "gin/model"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, kode int, message string) {
	var response m.ErrorResponse
	response.Status = kode //400 bad req, 404 not found, 500 internal server error, 401 unauthorized
	response.Message = message
}

func SendSuccesResponse(c *gin.Context, kode int, message string) {
	var response m.SuccessResponse
	response.Status = kode
	response.Message = message
}
