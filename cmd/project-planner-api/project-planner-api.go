package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/CuriousCrow/project-planner-service/cmd/project-planner-api/wrappers"
	"github.com/CuriousCrow/project-planner-service/configs"
	app_init "github.com/CuriousCrow/project-planner-service/init"
	project_planner_api "github.com/CuriousCrow/project-planner-service/internal/app/project-planner-api"
	_ "github.com/CuriousCrow/project-planner-service/internal/metrics"
	"github.com/CuriousCrow/project-planner-service/internal/services/projects"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/prometheus/client_golang/prometheus"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	appConfig, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	client := app_init.MongoClient(appConfig)
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Ошибка при отключении: %v", err)
		}
	}()
	db := client.Database(appConfig.Mongo.Name)

	app := gin.Default()

	// Эндпоинт, который будет вызывать Prometheus для сбора метрик
	//http.Handle("/metrics", promhttp.Handler())

	app.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	url := ginSwagger.URL("http://localhost:8081/docs/swagger.json")
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	projectSrv := projects.NewService(db)
	apiImpl := project_planner_api.NewImplementation(projectSrv)

	app.GET("/project/:id", wrappers.GetHandlerWrapper(apiImpl.FindProjectByID))
	app.PUT("/project/:id", wrappers.HandlerWrapper(apiImpl.UpdateProject))
	app.POST("/project/new", wrappers.HandlerWrapper(apiImpl.NewProjectFromTemplate))

	err = app.Run(fmt.Sprintf(":%d", appConfig.App.Port))
	if err != nil {
		panic(err)
	}
}
