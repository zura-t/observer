package server

import (
	"net/http"

	// fkng error because of go and gopls versions
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	diaryController "github.com/zura-t/observer.dev/internal/app/controller/http/diary"
	userController "github.com/zura-t/observer.dev/internal/app/controller/http/user"
	"github.com/zura-t/observer.dev/internal/app/controller/middleware"
	diaryUsecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
	userUsecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
	"github.com/zura-t/observer.dev/pkg/logger"
	"github.com/zura-t/observer.dev/pkg/token"
)

func NewRouter(handler *gin.Engine, userUsecase userUsecase.UserUsecase, diaryUsecase diaryUsecase.DiaryUsecase, tokenMaker token.Maker, logger logger.Interface) {
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	handler.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	authRoutes := handler.Group("/").Use(middleware.AuthMiddleware(tokenMaker))
	{
		userController.New(handler, userUsecase, tokenMaker, logger)
		diaryController.New(authRoutes, diaryUsecase, logger)
	}
}
