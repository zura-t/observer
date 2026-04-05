package notesController

import (
	"github.com/gin-gonic/gin"
	usecase "github.com/zura-t/observer.dev/internal/app/usecases/notes"
	"github.com/zura-t/observer.dev/pkg/logger"
)

type notesController struct {
	notesUsecase usecase.NotesUsecase
	logger       *logger.Logger
}

func New(handler gin.IRoutes, notesUsecase usecase.NotesUsecase, logger *logger.Logger) {
	routes := &notesController{notesUsecase, logger}
	_ = routes
}
