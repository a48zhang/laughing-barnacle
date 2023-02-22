package handler

import (
	"github.com/gin-gonic/gin"
	"main/service"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
} //@name Response

func SendResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "OK",
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, err error, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    403,
		Message: service.ErrorSender() + ": " + err.Error(),
		Data:    data,
	})
}

func SendError(c *gin.Context, err error, data interface{}, code int) {
	c.JSON(code, Response{
		Code:    code,
		Message: service.ErrorSender() + ": " + err.Error(),
		Data:    data,
	})
}
