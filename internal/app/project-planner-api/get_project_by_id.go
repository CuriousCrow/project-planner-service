package project_planner_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (impl *Implementation) FindProjectById(c *gin.Context) {
	ctx := c.Request.Context()

	hexID := c.Param("id")
	projectID, err := primitive.ObjectIDFromHex(hexID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	project, err := impl.service.FindProjectById(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, project)
}
