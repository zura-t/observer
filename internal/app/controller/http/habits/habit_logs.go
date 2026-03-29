package habitsController

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	controller "github.com/zura-t/observer.dev/internal/app/controller/http"
	"github.com/zura-t/observer.dev/internal/app/models"
	usecase "github.com/zura-t/observer.dev/internal/app/usecases/habits"
)

func (h *habitsController) createHabitLog(c *gin.Context) {
	var uri HabitIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		h.logger.Error(err, "habits routes - createHabitLog")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var req CreateHabitLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(err, "habits routes - createHabitLog")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	habitLog, err := h.habitsUsecase.CreateHabitLog(c, &usecase.CreateHabitLog{
		HabitID:     uri.HabitID,
		LogDate:     req.LogDate,
		ActualCount: req.ActualCount,
		Note:        req.Note,
	})
	if err != nil {
		h.logger.Error(err, "habits routes - createHabitLog")
		if errors.Is(err, models.ErrHabitNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, habitLog)
}

func (h *habitsController) getHabitLogs(c *gin.Context) {
	var uri HabitIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		h.logger.Error(err, "habits routes - getHabitLogs")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var req GetHabitLogsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		h.logger.Error(err, "habits routes - getHabitLogs")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabitLogs")
		return
	}

	logs, err := h.habitsUsecase.GetHabitLogs(c, &usecase.GetHabitLogsFilter{
		HabitID: uri.HabitID,
		UserID:  payload.ID,
		Limit:   req.Limit,
		Offset:  req.Offset,
	})
	if err != nil {
		h.logger.Error(err, "habits routes - getHabitLogs")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, logs)
}

func (h *habitsController) getHabitLog(c *gin.Context) {
	var req GetHabitLogByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.logger.Error(err, "habits routes - getHabitLog")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabitLog")
		return
	}

	log, err := h.habitsUsecase.GetHabitLog(c, req.ID, payload.ID)
	if err != nil {
		h.logger.Error(err, "habits routes - getHabitLog")
		if errors.Is(err, models.ErrHabitLogNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, log)
}

func (h *habitsController) updateHabitLog(c *gin.Context) {
	var uri GetHabitLogByIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		h.logger.Error(err, "habits routes - updateHabitLog")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var req UpdateHabitLogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error(err, "habits routes - updateHabitLog")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - updateHabitLog")
		return
	}

	habitLog, err := h.habitsUsecase.UpdateHabitLog(c, &usecase.UpdateHabitLog{
		ID:          uri.ID,
		LogDate:     req.LogDate,
		ActualCount: req.ActualCount,
		Note:        req.Note,
		UserID:      payload.ID,
	})
	if err != nil {
		h.logger.Error(err, "habits routes - updateHabitLog")
		if errors.Is(err, models.ErrHabitLogNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, habitLog)
}

func (h *habitsController) deleteHabitLog(c *gin.Context) {
	var req DeleteHabitLogRequest
	if err := c.ShouldBindUri(&req); err != nil {
		h.logger.Error(err, "habits routes - deleteHabitLog")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - deleteHabitLog")
		return
	}

	if err := h.habitsUsecase.DeleteHabitLog(c, req.ID, payload.ID); err != nil {
		h.logger.Error(err, "habits routes - deleteHabitLog")
		if errors.Is(err, models.ErrHabitLogNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "habit log deleted"})
}

func (h *habitsController) deleteHabitLogs(c *gin.Context) {
	var uri HabitIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		h.logger.Error(err, "habits routes - deleteHabitLogs")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - deleteHabitLogs")
		return
	}

	if err := h.habitsUsecase.DeleteHabitLogs(c, uri.HabitID, payload.ID); err != nil {
		h.logger.Error(err, "habits routes - deleteHabitLogs")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "habit logs deleted"})
}

func (h *habitsController) deleteAllHabitLogs(c *gin.Context) {
	payload, err := controller.GetPayload(c)
	if err != nil {
		h.logger.Error(err, "habits routes - deleteAllHabitLogs")
		return
	}

	if err := h.habitsUsecase.DeleteAllHabitLogs(c, payload.ID); err != nil {
		h.logger.Error(err, "habits routes - deleteAllHabitLogs")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "all habit logs deleted"})
}
