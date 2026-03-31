package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	gotodo "github.com/grancc/go-to-do-app"
)

// SignUp registers a new user.
// @Summary Register
// @Description Create account; password is stored hashed server-side
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body gotodo.User true "User profile"
// @Success 200 {object} IdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input gotodo.User

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, IdResponse{Id: id})
}

// SignInInput is login credentials (plain password field is JSON key password_hash for compatibility).
type SignInInput struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password_hash" binding:"required"`
}

// SignIn returns JWT for subsequent requests (Authorization: Bearer <token>).
// @Summary Sign in
// @Tags Auth
// @Accept json
// @Produce json
// @Param input body SignInInput true "Credentials"
// @Success 200 {object} TokenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.UserName, input.Password)

	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, TokenResponse{Token: token})
}
