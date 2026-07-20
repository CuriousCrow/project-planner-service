package projects

import (
	"context"
	"fmt"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
)

func (s service) UpdateProject(ctx context.Context, project dto.Project) error {

	updateResult, err := s.db.Collection("projects").UpdateByID(ctx, "$"+project.ID.Hex(), project)
	if err != nil {
		return fmt.Errorf("error updating project %s: %w", project.ID, err)
	}
	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("project %s not found", project.ID)
	}

	return nil
}
