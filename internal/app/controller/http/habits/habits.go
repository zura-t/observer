package habitsController

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/zura-t/observer.dev/internal/app/controller/http"
	"github.com/zura-t/observer.dev/internal/app/models"
	usecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
	"github.com/zura-t/observer.dev/pkg/logger"
)

type habitsController struct {
	habitsUsecase usecase.HabitsUsecase
	logger        logger.Interface
}

func New(handler gin.IRoutes, habitsUsecase usecase.HabitsUsecase, logger logger.Interface) {
	routes := &habitsController{habitsUsecase, logger}

	handler.POST("/habits", routes.createHabit)
	handler.GET("/habits", routes.getHabits)
	handler.GET("/habits/:id", routes.getHabit)
	handler.PATCH("/habits/:id", routes.updateHabit)
	handler.DELETE("/habits/:id", routes.deleteHabit)

	handler.POST("/habit_logs/:habitID", routes.createHabitLog)
	handler.GET("/habit_logs/by_habit/:habitID", routes.getHabitLogs)
	handler.GET("/habit_logs/:id", routes.getHabitLog)
	handler.PATCH("/habit_logs/:id", routes.updateHabitLog)
	handler.DELETE("/habit_logs/:id", routes.deleteHabitLog)
	handler.DELETE("/habit_logs/by_habit/:habitID", routes.deleteHabitLogs)
	handler.DELETE("/habit_logs", routes.deleteAllHabitLogs)
}

func (h *habitsController) createHabit(c *gin.Context) {
	var req CreateHabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(err, "habits routes - createHabit")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - createHabit")
		return
	}

	habit, err := h.habitsUsecase.CreateHabit(c, &usecase.CreateHabit{
		Title:       req.Title,
		Frequency:   req.Frequency,
		TargetCount: req.TargetCount,
		UserID:      payload.ID,
	})
	if err != nil {
		h.logger.Error(err, "habits routes - createHabit")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, habit)
}

func (h *habitsController) getHabits(c *gin.Context) {
	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabits")
		return
	}

	habits, err := h.habitsUsecase.GetHabits(c, payload.ID)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabits")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, habits)
}

func (h *habitsController) getHabit(c *gin.Context) {
	var req GetHabitByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.logger.Error(err, "habits routes - getHabit")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabit")
		return
	}

	habit, err := h.habitsUsecase.GetHabit(c, req.ID, payload.ID)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabit")
		if errors.Is(err, models.ErrHabitNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, habit)
}

func (h *habitsController) updateHabit(c *gin.Context) {
	var uri GetHabitByIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		h.logger.Error(err, "habits routes - updateHabit")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var req UpdateHabitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(err, "habits routes - updateHabit")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - updateHabit")
		return
	}

	habit, err := h.habitsUsecase.UpdateHabit(c, &usecase.UpdateHabit{
		ID:          uri.ID,
		Title:       req.Title,
		Frequency:   req.Frequency,
		TargetCount: req.TargetCount,
		UserID:      payload.ID,
	})
	if err != nil {
		h.logger.Error(err, "habits routes - updateHabit")
		if errors.Is(err, models.ErrHabitNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, habit)
}

func (h *habitsController) deleteHabit(c *gin.Context) {
	var req DeleteHabitRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.logger.Error(err, "habits routes - deleteHabit")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - deleteHabit")
		return
	}

	if err := h.habitsUsecase.DeleteHabit(c, req.ID, payload.ID); err != nil {
		h.logger.Error(err, "habits routes - deleteHabit")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "habit deleted"})
}
