package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	gotodo "github.com/grancc/go-to-do-app"
)

// CreateItem adds a todo item to a list.
// @Summary Create item
// @Tags Items
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param id path int true "List id"
// @Param input body gotodo.ToDoItem true "New item"
// @Success 200 {object} IdResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/{id}/items [post]
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

	c.JSON(http.StatusOK, IdResponse{Id: id})
}

type getAllItemsResponse struct {
	Data []gotodo.ToDoItem `json:"data"`
}

// GetAllItems returns items for a list.
// @Summary List items in a list
// @Tags Items
// @Security ApiKeyAuth
// @Produce json
// @Param id path int true "List id"
// @Success 200 {object} getAllItemsResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/{id}/items [get]
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

// GetItemById returns a single item for the user.
// @Summary Get item by id
// @Tags Items
// @Security ApiKeyAuth
// @Produce json
// @Param item_id path int true "Item id"
// @Success 200 {object} gotodo.ToDoItem
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/items/{item_id} [get]
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

// @Summary Update item
// @Tags Items
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param item_id path int true "Item id"
// @Param input body gotodo.UpdateListItemInput true "Fields to update"
// @Success 200 {object} StatusResponse
// @Failure 400 {object} ErrorResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/items/{item_id} [put]
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

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}

// DeleteItem removes an item.
// @Summary Delete item
// @Tags Items
// @Security ApiKeyAuth
// @Produce json
// @Param item_id path int true "Item id"
// @Success 200 {object} StatusResponse
// @Failure 401 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/items/{item_id} [delete]
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

	c.JSON(http.StatusOK, StatusResponse{
		Status: "ok",
	})
}
