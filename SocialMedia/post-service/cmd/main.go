package main

import (
	"log"
	"net/http"

	"github.com/rguarnizo/SocialMedia/pkg/config"
	"github.com/rguarnizo/SocialMedia/post-service/internal/handler"
	"github.com/rguarnizo/SocialMedia/post-service/internal/router"

	"github.com/rguarnizo/SocialMedia/post-service/internal/repository"
	"github.com/rguarnizo/SocialMedia/post-service/internal/service"
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

	repo := repository.NewPostRepository(db)
	svc := service.NewPostService(repo)
	postsHandler := handler.NewPostHandler(svc)

	router := router.SetupRouter(cfg.JWTSecret, postsHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
