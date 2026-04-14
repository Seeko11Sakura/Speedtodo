package task

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	items []Task
}

func NewHandler() *Handler {
	return &Handler{items: []Task{}}
}

func (h *Handler) Create(c *gin.Context) {
	var input Task
	if err := c.ShouldBindJSON(&input); err != nil || input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_input", "message": "title is required"})
		return
	}
	input.ID = uint(len(h.items) + 1)
	input.Status = "active"
	h.items = append(h.items, input)
	c.JSON(http.StatusCreated, input)
}

func (h *Handler) List(c *gin.Context) {
	c.JSON(http.StatusOK, h.items)
}

func (h *Handler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid task id"})
		return
	}
	for _, item := range h.items {
		if item.ID == uint(id) {
			c.JSON(http.StatusOK, item)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "task not found"})
}

func (h *Handler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid task id"})
		return
	}
	var input Task
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
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "task not found"})
}

func (h *Handler) Complete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid task id"})
		return
	}
	for i, item := range h.items {
		if item.ID == uint(id) {
			h.items[i].Status = "completed"
			now := time.Now().UTC()
			h.items[i].CompletedAt = &now
			c.JSON(http.StatusOK, h.items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "task not found"})
}

func (h *Handler) Reopen(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid task id"})
		return
	}
	for i, item := range h.items {
		if item.ID == uint(id) {
			h.items[i].Status = "active"
			h.items[i].CompletedAt = nil
			c.JSON(http.StatusOK, h.items[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "task not found"})
}

func (h *Handler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "invalid_id", "message": "invalid task id"})
		return
	}
	for i, item := range h.items {
		if item.ID == uint(id) {
			h.items = append(h.items[:i], h.items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"status": "deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"code": "not_found", "message": "task not found"})
}
