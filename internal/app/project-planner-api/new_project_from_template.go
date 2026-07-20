package project_planner_api

import (
	"context"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"github.com/CuriousCrow/project-planner-service/typed_error"
	"github.com/gin-gonic/gin"
)

// NewProjectFromTemplate ...
func (impl *Implementation) NewProjectFromTemplate(ctx context.Context, _ gin.Params, _ EmptyRequest) (*dto.Project, error) {

	newProject, err := impl.service.NewProjectFromTemplate(ctx)
	if err != nil {
		return &dto.Project{}, typed_error.NewServerError(err)
	}

	return &newProject, nil
}
