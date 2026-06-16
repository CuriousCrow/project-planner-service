package project_planner_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// FindProjectByID ...
func (impl *Implementation) FindProjectByID(c *gin.Context) {
	ctx := c.Request.Context()

	hexID := c.Param("id")
	projectID, err := bson.ObjectIDFromHex(hexID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	project, err := impl.service.FindProjectByID(ctx, projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, project)
}
