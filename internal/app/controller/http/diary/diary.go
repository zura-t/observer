package diaryController

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	controller "github.com/zura-t/observer.dev/internal/app/controller/http"
	"github.com/zura-t/observer.dev/internal/app/models"
	usecase "github.com/zura-t/observer.dev/internal/app/usecases/diary"
	"github.com/zura-t/observer.dev/pkg/logger"
)

type diaryController struct {
	diaryUsecase usecase.DiaryUsecase
	logger       logger.Interface
}

func New(handler gin.IRoutes, diaryUsecase usecase.DiaryUsecase, logger logger.Interface) {
	routes := &diaryController{diaryUsecase, logger}

	handler.POST("/diary", routes.createDiary)
	handler.GET("/diary/:id", routes.getDiaryByID)
	handler.GET("/diary", routes.getDiary)
	handler.PATCH("/diary/:id", routes.updateDiary)
	handler.DELETE("/diary/:id", routes.deleteDiary)
}

func (d *diaryController) createDiary(c *gin.Context) {
	var req CreateDiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		d.logger.Error(err, "diary routes - createDiary")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		d.logger.Error(err, "diary routes - createDiary")
		return
	}

	entry, err := d.diaryUsecase.CreateDiaryEntry(c, &usecase.CreateDiaryEntry{
		Title:     req.Title,
		Text:      req.Text,
		EntryDate: time.Now(),
		UserID:    payload.ID,
	})
	if err != nil {
		d.logger.Error(err, "diary routes - createDiary")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, entry)
}

func (d *diaryController) getDiaryByID(c *gin.Context) {
	var req GetDiaryByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		d.logger.Error(err, "diary routes - getDiaryByID")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	// ? check user id

	entry, err := d.diaryUsecase.GetEntry(c, req.ID)
	if err != nil {
		d.logger.Error(err, "diary routes - getDiaryByID")
		if errors.Is(err, models.ErrDiaryEntryNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, entry)
}

func (d *diaryController) getDiary(c *gin.Context) {
	var req GetDiaryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		d.logger.Error(err, "diary routes - getDiary")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	payload, err := controller.GetPayload(c)
	if err != nil {
		d.logger.Error(err, "diary routes - getDiary")
		return
	}

	filter := &usecase.DiarySearchFilter{
		UserID: payload.ID,
	}

	if req.DateFrom != "" {
		dateFrom, err := time.Parse("2000-01-01", req.DateFrom)
		if err != nil {
			d.logger.Error(err, "diary routes - getDiary")
			controller.ErrorResponse(c, http.StatusBadRequest, errors.New("invalid date_from format, expected YYYY-MM-DD"))
			return
		}
		filter.DateFrom = &dateFrom
	}

	if req.DateTo != "" {
		dateTo, err := time.Parse("2000-01-01", req.DateTo)
		if err != nil {
			d.logger.Error(err, "diary routes - getDiary")
			controller.ErrorResponse(c, http.StatusBadRequest, errors.New("invalid date_to format, expected YYYY-MM-DD"))
			return
		}
		filter.DateTo = &dateTo
	}

	entries, err := d.diaryUsecase.GetEntries(c, filter)
	if err != nil {
		d.logger.Error(err, "diary routes - getDiary")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, entries)
}

func (d *diaryController) updateDiary(c *gin.Context) {
	var uri GetDiaryByIDRequest
	if err := c.ShouldBindUri(&uri); err != nil {
		d.logger.Error(err, "diary routes - updateDiary")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	var req UpdateDiaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		d.logger.Error(err, "diary routes - updateDiary")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	update := &usecase.UpdateDiaryEntry{
		ID:    uri.ID,
		Title: req.Title,
		Text:  req.Text,
	}

	if req.EntryDate != "" {
		entryDate, err := time.Parse("2000-01-01", req.EntryDate)
		if err != nil {
			d.logger.Error(err, "diary routes - updateDiary")
			controller.ErrorResponse(c, http.StatusBadRequest, errors.New("invalid entry_date format, expected YYYY-MM-DD"))
			return
		}
		update.EntryDate = &entryDate
	}

	diaryUpdated, err := d.diaryUsecase.UpdateDiaryEntry(c, update)
	if err != nil {
		d.logger.Error(err, "diary routes - updateDiary")
		if errors.Is(err, models.ErrDiaryEntryNotFound) {
			controller.ErrorResponse(c, http.StatusNotFound, err)
			return
		}
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, diaryUpdated)
}

func (d *diaryController) deleteDiary(c *gin.Context) {
	var req DeleteDiaryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		d.logger.Error(err, "diary routes - deleteDiary")
		controller.ErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if err := d.diaryUsecase.DeleteEntry(c, req.ID); err != nil {
		d.logger.Error(err, "diary routes - deleteDiary")
		controller.ErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "entry deleted"})
}
