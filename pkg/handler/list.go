package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gotodo "github.com/grancc/go-to-do-app"
)

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

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllListsResponse struct {
	Data []gotodo.ToDoList `json:"data"`
}

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

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

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

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
