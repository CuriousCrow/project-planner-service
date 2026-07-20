package project_planner_api

import (
	"context"

	"github.com/CuriousCrow/project-planner-service/internal/dto"
	"github.com/CuriousCrow/project-planner-service/typed_error"
	"github.com/gin-gonic/gin"
)

func (impl *Implementation) UpdateProject(ctx context.Context, _ gin.Params, project dto.Project) (*EmptyResponse, error) {
	//ctx := c.Request.Context()
	//
	//bodyBytes := make([]byte, c.Request.ContentLength)
	//
	//n, err := c.Request.Body.Read(bodyBytes)
	//if err != nil && err.Error() != "EOF" {
	//	c.JSON(400, gin.H{"error": err.Error()})
	//}
	//fmt.Println("read request body: ", string(bodyBytes[:n]))
	//
	//var project dto.Project
	//
	//err = json.Unmarshal(bodyBytes, &project)
	//if err != nil {
	//	c.JSON(400, gin.H{"error": err.Error()})
	//	return
	//}

	err := impl.service.UpdateProject(ctx, project)
	if err != nil {
		return &EmptyResponse{}, typed_error.NewServerError(err)
	}

	return &EmptyResponse{}, nil
}
