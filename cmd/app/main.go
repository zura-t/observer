package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	diaryRepo "github.com/zura-t/observer.dev/internal/app/repositories/diary"
	userRepo "github.com/zura-t/observer.dev/internal/app/repositories/user"

	"github.com/zura-t/observer.dev/internal/app/server"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
	userUsecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
	"github.com/zura-t/observer.dev/internal/config"
	"github.com/zura-t/observer.dev/pkg/logger"
	"github.com/zura-t/observer.dev/pkg/token"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("can't load config ", err)
	}

	l := logger.New(config.LogLevel)
	conn, err := pgx.Connect(context.Background(), "postgresql://postgres:root@host.docker.internal:5423/observer?sslmode=disable")
	if err != nil {
		log.Fatal("Error ", err)
	}
	defer conn.Close(context.Background())

	//! repos
	userRepo := userRepo.New(conn)
	diaryRepo := diaryRepo.New(conn)

	tokenMaker, err := token.NewJwtMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("Error creating token maker: ", err)
	}

	//! usecases
	userUsecase := userUsecase.New(userRepo, tokenMaker, config)
	diaryUsecase := diaryUsecase.New(diaryRepo, config)

	handler := gin.New()
	server.NewRouter(handler, userUsecase, diaryUsecase, tokenMaker, l)
	handler.Run("127.0.0.1:8080")
}
