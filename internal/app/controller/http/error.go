package controller

import (
	"errors"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zura-t/observer.dev/internal/app/models"
)

func ErrorResponse(c *gin.Context, code int, err error) {
	if err == io.EOF {
		err = errors.New("Empty body")
	}
	if err == models.ErrEmailAlreadyExists {
		code = http.StatusConflict
	}
	c.AbortWithStatusJSON(code, gin.H{"error": err.Error()})
	return
}
