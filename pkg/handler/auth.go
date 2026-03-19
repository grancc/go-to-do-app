package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gotodo "github.com/grancc/go-to-do-app"
)

func (h *Handler) signUp(c *gin.Context) {
	var input gotodo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {

}
