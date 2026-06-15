package project_planner_api

import (
	"context"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProjectService interface {
	FindProjectById(ctx context.Context, projectID primitive.ObjectID) (dto.Project, error)
	NewProjectFromTemplate(ctx context.Context) (dto.Project, error)
}
