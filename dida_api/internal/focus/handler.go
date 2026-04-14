package focus

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	items []Session
}

func NewHandler() *Handler {
	return &Handler{items: []Session{}}
}

func (h *Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, h.items)
}

func (h *Handler) Create(c *gin.Context) {
	var input Session
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_input", "message": "invalid request body"})
		return
	}
	input.ID = uint(len(h.items) + 1)
	input.Status = "running"
	h.items = append(h.items, input)
	c.JSON(http.StatusCreated, input)
}

func (h *Handler) Complete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid session id"})
		return
	}
	for i, item := range h.items {
		if item.ID == uint(id) {
			h.items[i].Complete()
			c.JSON(http.StatusOK, h.items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "session not found"})
}

func (h *Handler) Cancel(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid session id"})
		return
	}
	for i, item := range h.items {
		if item.ID == uint(id) {
			h.items[i].Cancel()
			c.JSON(http.StatusOK, h.items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "session not found"})
}
