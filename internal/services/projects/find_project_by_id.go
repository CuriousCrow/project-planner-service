package projects

import (
	"context"
	"fmt"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func (s service) FindProjectById(ctx context.Context, projectID primitive.ObjectID) (dto.Project, error) {
	var resProject dto.Project
	err := s.db.Collection("projects").FindOne(ctx, bson.M{"_id": projectID}).Decode(&resProject)
	if err != nil {
		return resProject, fmt.Errorf("projects.FindOne error: %w", err)
	}

	return resProject, nil
}
