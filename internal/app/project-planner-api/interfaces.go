package project_planner_api

import (
	"context"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// ProjectService ...
type ProjectService interface {
	FindProjectByID(ctx context.Context, projectID bson.ObjectID) (dto.Project, error)
	NewProjectFromTemplate(ctx context.Context) (dto.Project, error)
	UpdateProject(ctx context.Context, project dto.Project) error
}
