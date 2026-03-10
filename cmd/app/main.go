package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	userRepo "github.com/zura-t/observer.dev/internal/app/repositories/user"
	"github.com/zura-t/observer.dev/internal/app/server"
	userUsecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
	"github.com/zura-t/observer.dev/internal/config"
	"github.com/zura-t/observer.dev/pkg/logger"
	"github.com/zura-t/observer.dev/pkg/token"
)

func main() {
	config, err := config.LoadConfig("../../")
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

	tokenMaker, err := token.NewJwtMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal("Error creating token maker: ", err)
	}

	//! usecases
	userUsecase := userUsecase.New(userRepo, tokenMaker, config)

	handler := gin.New()
	server.NewRouter(handler, userUsecase, tokenMaker, l)
	handler.Run("127.0.0.1:8080")
}
