package handler

import (
	"fmt"
	"github.com/AndreCDiniz/ApiRestFull-Go/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(context *gin.Context, code int, msg string) {
	context.Header("Content-type", "application/json")
	context.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(context *gin.Context, op string, data interface{}) {
	context.Header("Content-type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}
