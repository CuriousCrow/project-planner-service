package project_planner_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (impl *Implementation) NewProjectFromTemplate(c *gin.Context) {
	ctx := c.Request.Context()

	newProject, err := impl.service.NewProjectFromTemplate(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newProject)
}
