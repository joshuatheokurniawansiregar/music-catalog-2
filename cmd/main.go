package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/configs"
	memberships_handler "github.com/joshuatheokurniawansiregar/music_catalog_2/internal/handler/memberships"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	memberships_repository "github.com/joshuatheokurniawansiregar/music_catalog_2/internal/repository/memberships"
	memberships_service "github.com/joshuatheokurniawansiregar/music_catalog_2/internal/service/memberships"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/pkg/internalsql"
)

func main() {
	
	var(
		cfg *configs.Config
	)

	var err error = configs.Init(
		configs.WithConfigFolder([]string{
			"./internal/configs",
		}),
		configs.WithConfigFile(
			"config",
		),
		configs.WithConfigType(
			"yaml",
		),
	)

	if err !=nil{
		log.Fatal("failed to initialize conifgs", err)
	}

	cfg = configs.GetConfig()

	database, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil{
		log.Fatalf("failed to connect database %v", err)
	}

	database.AutoMigrate(&memberships.User{})

	var engine *gin.Engine = gin.Default()
	var membershipsRepository *memberships_repository.Repository = memberships_repository.NewRepository(database)
	var membershipsService *memberships_service.Service = memberships_service.NewService(cfg, membershipsRepository)
	var membershipsHandler *memberships_handler.Handler = memberships_handler.NewHandler(engine, membershipsService)
	membershipsHandler.RegisterRoute()
	
	engine.Use(gin.Recovery())
	engine.Run(cfg.Service.Port)

}