package project_planner_api

type Implementation struct {
	service ProjectService
}

func NewImplementation(projectService ProjectService) *Implementation {
	return &Implementation{
		service: projectService,
	}
}
