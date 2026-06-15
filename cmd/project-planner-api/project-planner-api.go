package main

import (
	"context"
	"log"
	"time"

	app_init "github.com/CuriousCrow/project-planner-service/init"
	project_planner_api "github.com/CuriousCrow/project-planner-service/internal/app/project-planner-api"
	"github.com/CuriousCrow/project-planner-service/internal/services/projects"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := app_init.MongoClient(ctx)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Ошибка при отключении: %v", err)
		}
	}()
	db := client.Database(app_init.MONGO_DB)

	app := gin.Default()

	projectSrv := projects.NewService(db)
	apiImpl := project_planner_api.NewImplementation(projectSrv)
	app.GET("/project/:id", apiImpl.FindProjectById)
	app.POST("/project/new", apiImpl.NewProjectFromTemplate)

	app.Run(":8080")
}
