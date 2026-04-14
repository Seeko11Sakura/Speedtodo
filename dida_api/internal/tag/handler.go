package tag

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	items []Tag
}

func NewHandler() *Handler {
	return &Handler{items: []Tag{}}
}

func (h *Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, h.items)
}

func (h *Handler) Create(c *gin.Context) {
	var input Tag
	if err := c.ShouldBindJSON(&input); err != nil || input.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_input", "message": "name is required"})
		return
	}
	input.ID = uint(len(h.items) + 1)
	h.items = append(h.items, input)
	c.JSON(http.StatusCreated, input)
}
