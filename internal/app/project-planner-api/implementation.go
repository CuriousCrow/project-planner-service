package project_planner_api

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
