package projects

import (
	"context"
	"fmt"
	"log"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// NewProjectFromTemplate ...
func (s service) NewProjectFromTemplate(ctx context.Context) (dto.Project, error) {
	newProject := dto.ProjectTemplate
	newProject.ID = bson.NewObjectID()
	newProject.Title = "project " + newProject.ID.Hex()

	insertResult, err := s.db.Collection("projects").InsertOne(ctx, newProject)
	if err != nil {
		return dto.Project{}, fmt.Errorf("error inserting new project: %w", err)
	}

	log.Println(newProject.ID.Hex(), insertResult.InsertedID)

	return newProject, nil
}
