package main

import (
	"log"
	"net/http"

	"github.com/rguarnizo/SocialMedia/auth-service/internal/handler"
	"github.com/rguarnizo/SocialMedia/auth-service/internal/router"
	"github.com/rguarnizo/SocialMedia/pkg/config"

	"github.com/rguarnizo/SocialMedia/auth-service/internal/repository"
	"github.com/rguarnizo/SocialMedia/auth-service/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate()

	repo := repository.NewUserRepository(db)
	svc := service.NewAuthService(repo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(svc)

	router := router.SetupRouter(authHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
