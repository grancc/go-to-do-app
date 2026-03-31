package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// StatusResponse confirms a mutation succeeded.
type StatusResponse struct {
	Status string `json:"status"`
}

// IdResponse returns created entity id.
type IdResponse struct {
	Id int `json:"id"`
}

// TokenResponse returns JWT after sign-in.
type TokenResponse struct {
	Token string `json:"token"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
