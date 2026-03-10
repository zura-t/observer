package userController

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/zura-t/observer.dev/internal/app/controller/http"
	"github.com/zura-t/observer.dev/internal/app/controller/middleware"
	"github.com/zura-t/observer.dev/internal/app/models"
	usecase "github.com/zura-t/observer.dev/internal/app/usecases/user"
	"github.com/zura-t/observer.dev/pkg/logger"
	"github.com/zura-t/observer.dev/pkg/token"
)

type userController struct {
	userUsecase usecase.UserUsecase
	logger      logger.Interface
}

func New(handler *gin.Engine, userUsecase usecase.UserUsecase, tokenMaker token.Maker, logger logger.Interface) {
	routes := &userController{userUsecase, logger}

	handler.POST("/register", routes.createUser)
	handler.POST("/login", routes.login)
	handler.POST("/renew", routes.renewAccessToken)
	handler.POST("/logout", routes.logout)

	authRoutes := handler.Group("/").Use(middleware.AuthMiddleware(tokenMaker))
	authRoutes.GET("/users/:id", routes.getUserByID)
	authRoutes.GET("/users/email", routes.getUserByEmail)
	authRoutes.PATCH("/users/:id", routes.updateUser)
	authRoutes.DELETE("/users/:id", routes.deleteUser)
}

func (u *userController) createUser(c *gin.Context) {
	var req controller.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Error(err, "user routes - createUser")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}
	user, err := u.userUsecase.RegisterUser(c, &usecase.RegisterUser{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		u.logger.Error(err, "user routes - createUser")
		if errors.Is(err, models.ErrEmailAlreadyExists) {
			controller.ErrorResponse(c, http.StatusConflict, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("refresh_token", user.RefreshToken, int(time.Until(user.RefreshTokenExpiresAt).Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusCreated, user)
}

func (u *userController) login(c *gin.Context) {
	var req controller.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Error(err, "user routes - login")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.Login(c, &usecase.Login{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		u.logger.Error(err, "user routes - login")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.SetCookie("refresh_token", user.RefreshToken, int(time.Until(user.RefreshTokenExpiresAt).Seconds()), "/", "localhost", false, true)

	c.JSON(http.StatusOK, user)
}

func (u *userController) getUserByID(c *gin.Context) {
	var req controller.GetUserByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		u.logger.Error(err, "user routes - getUserByID")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.GetUserByID(c, req.ID)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		u.logger.Error(err, "user routes - getUserByID")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) getUserByEmail(c *gin.Context) {
	var req controller.GetUserByEmailRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		u.logger.Error(err, "user routes - getUserByEmail")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.GetUserByEmail(c, req.Email)
	if err != nil {
		u.logger.Error(err, "user routes - getUserByEmail")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) updateUser(c *gin.Context) {
	var uri controller.GetUserByIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		u.logger.Error(err, "user routes - updateUser")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var req controller.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.logger.Error(err, "user routes - updateUser")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.UpdateUser(c, uri.ID, &usecase.UpdateUser{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		u.logger.Error(err, "user routes - updateUser")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) deleteUser(c *gin.Context) {
	var req controller.DeleteUserRequest
	if err := c.ShouldBindUri(&req); err != nil {
		u.logger.Error(err, "user routes - deleteUser")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	err := u.userUsecase.DeleteUser(c, req.ID)
	if err != nil {
		u.logger.Error(err, "user routes - deleteUser")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (u *userController) renewAccessToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		u.logger.Error(err, "user routes - renewAccessToken")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := u.userUsecase.RenewAccessToken(c, refreshToken)
	if err != nil {
		u.logger.Error(err, "user routes - renewAccessToken")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (u *userController) logout(c *gin.Context) {
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
