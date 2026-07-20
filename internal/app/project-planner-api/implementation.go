package project_planner_api

// EmptyRequest ...
type EmptyRequest struct {
}

// EmptyResponse ...
type EmptyResponse struct{}

// Implementation ...
type Implementation struct {
	service ProjectService
}

// NewImplementation ...
func NewImplementation(projectService ProjectService) *Implementation {
	return &Implementation{
		service: projectService,
	}
}
