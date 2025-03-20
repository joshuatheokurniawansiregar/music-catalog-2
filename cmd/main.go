package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/configs"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/pkg/internalsql"
)

func main() {
	r := gin.Default()
	var(
		cfg *configs.Config
	)

	var err error = configs.Init(
		configs.WithConfigFolder([]string{
			"./internal/con",
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
	


}