package list

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	items []List
}

func NewHandler() *Handler {
	return &Handler{items: []List{}}
}

func (h *Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, h.items)
}

func (h *Handler) Create(c *gin.Context) {
	var input List
	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_input", "message": "name is required"})
		return
	}
	input.ID = uint(len(h.items) + 1)
	h.items = append(h.items, input)
	c.JSON(http.StatusCreated, input)
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid list id"})
		return
	}
	var input List
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_input", "message": "invalid request body"})
		return
	}
	for i, item := range h.items {
		if item.ID == uint(id) {
			input.ID = uint(id)
			h.items[i] = input
			c.JSON(http.StatusOK, input)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "list not found"})
}
