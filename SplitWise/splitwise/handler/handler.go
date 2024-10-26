package handler

import (
	"net/http"
	"splitwise/models"
	"splitwise/service"

	"github.com/gin-gonic/gin"
)

// *This layer handles the HTTP requests and delegates the work to the appropriate service layer.

type GroupHandler struct {
	service service.GroupService
}

func NewGroupHandler(uc service.GroupService) *GroupHandler {
	return &GroupHandler{
		service: uc,
	}
}

func (h *GroupHandler) CreateGroup(c *gin.Context) {
	var group models.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.service.CreateGroup(group); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create group"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Group created successfully"})
}

func (h *GroupHandler) GetGroups(c *gin.Context) {
	groups, err := h.service.GetGroups()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get groups"})
		return
	}

	c.JSON(http.StatusOK, groups)
}
