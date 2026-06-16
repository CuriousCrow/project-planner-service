package projects

import (
	"context"
	"fmt"

	"github.com/CuriousCrow/project-planner-service/internal/db"
	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// FindProjectByID ...
func (s service) FindProjectByID(ctx context.Context, projectID bson.ObjectID) (dto.Project, error) {
	var resProject dto.Project
	err := s.db.Collection(db.CollectionProjects).FindOne(ctx, bson.M{"_id": projectID}).Decode(&resProject)
	if err != nil {
		return resProject, fmt.Errorf("projects.FindOne error: %w", err)
	}

	return resProject, nil
}
