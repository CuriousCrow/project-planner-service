package project_planner_api

import (
	"context"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"github.com/CuriousCrow/project-planner-service/typed_error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// FindProjectByID ...
func (impl *Implementation) FindProjectByID(ctx context.Context, params gin.Params, _ EmptyRequest) (*dto.Project, error) {

	hexID, _ := params.Get("id")
	projectID, err := bson.ObjectIDFromHex(hexID)
	if err != nil {
		return &dto.Project{}, typed_error.New(typed_error.BadRequest, "invalid id")
	}

	project, err := impl.service.FindProjectByID(ctx, projectID)
	if err != nil {
		return &dto.Project{}, typed_error.New(typed_error.ServerError, err.Error())
	}

	return &project, nil
}
