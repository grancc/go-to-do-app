package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gotodo "github.com/grancc/go-to-do-app"
)

// CreateList adds a todo list for the authenticated user.
// @Summary Create list
// @Tags Lists
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param input body gotodo.ToDoList true "New list"
// @Success 200 {object} IdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	var input gotodo.ToDoList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.services.TodoList.Create(id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, IdResponse{Id: id})
}

type getAllListsResponse struct {
	Data []gotodo.ToDoList `json:"data"`
}

// GetLists returns all lists of the current user.
// @Summary List all todo lists
// @Tags Lists
// @Security ApiKeyAuth
// @Produce json
// @Success 200 {object} getAllListsResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists [get]
func (h *Handler) getList(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	lists, err := h.services.TodoList.GetAll(id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}

// GetListById returns one list by id if it belongs to the user.
// @Summary Get list by id
// @Tags Lists
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "List id"
// @Success 200 {object} gotodo.ToDoList
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	list, err := h.services.TodoList.GetById(userid, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// UpdateList patches title/description.
// @Summary Update list
// @Tags Lists
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "List id"
// @Param input body gotodo.UpdateListInput true "Fields to update"
// @Success 200 {object} StatusResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	var input gotodo.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.UpdateList(userid, id, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// DeleteList removes a list owned by the user.
// @Summary Delete list
// @Tags Lists
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "List id"
// @Success 200 {object} StatusResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userid, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userid, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
