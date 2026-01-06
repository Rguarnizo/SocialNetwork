package main

import (
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rguarnizo/SocialMedia/pkg/config"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/handler"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/repository"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/router"
	"github.com/rguarnizo/SocialMedia/profile-service/internal/service"
)

func main() {

	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewUserRepository(db)
	svc := service.NewUserService(repo)
	profileHandler := handler.NewUserHandler(svc)

	router := router.SetupRouter(cfg.JWTSecret, profileHandler)

	log.Println("User Service running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
