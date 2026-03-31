package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gotodo "github.com/grancc/go-to-do-app"
)

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}

	var input gotodo.ToDoItem
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.services.ToodItem.Create(userId, listId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllItemsResponse struct {
	Data []gotodo.ToDoItem `json:"data"`
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid list id param")
		return
	}

	items, err := h.services.ToodItem.GetAllItems(userId, listId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllItemsResponse{
		Data: items,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid item id param")
		return
	}

	item, err := h.services.ToodItem.GetItemById(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid item id param")
		return
	}

	var input gotodo.UpdateListItemInput
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.ToodItem.UpdateItem(userId, itemId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	itemId, err := strconv.Atoi(c.Param("item_id"))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, "invalid item id param")
		return
	}

	err = h.services.ToodItem.DeleteItem(userId, itemId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
