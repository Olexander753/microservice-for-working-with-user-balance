package main

import (
	"log"

	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/config"
	"github.com/Olexander753/microservice-for-working-with-user-balance/internal/server"

	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/handler"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/repository"
	"github.com/Olexander753/microservice-for-working-with-user-balance/pkg/service"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// read confid file
	gin.SetMode(gin.ReleaseMode)
	cfg := config.GetConfig()

	//connect to postgresql db
	db, err := repository.NewPostgresBD(cfg)
	if err != nil {
		log.Fatal("Failed connect to postgres, error: ", err)
	}

	//create repository
	repo := repository.NewRepository(db)

	//create service
	services := service.NewService(repo)

	//create handler
	handlers := handler.NewHandler(services)

	//create http router
	srv := new(server.Server)
	err = srv.Run(cfg, handlers.InutRoutes())
	if err != nil {
		log.Fatal("error http server: ", err)
	}
}
