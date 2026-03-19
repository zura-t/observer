package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zura-t/observer.dev/internal/app/controller/middleware"
	"github.com/zura-t/observer.dev/pkg/token"
)

func GetPayload(c *gin.Context) (*token.Payload, error) {
	var payload *token.Payload
	payloadData, exists := c.Get(middleware.AuthorizationPayloadKey)
	if !exists {
		err := errors.New("authorization payload not found")
		ErrorResponse(c, http.StatusInternalServerError, err)
		return nil, err
	}
	data, ok := payloadData.(*token.Payload)
	if ok {
		payload = data
	} else {
		err := errors.New("invalid authorization payload")
		ErrorResponse(c, http.StatusInternalServerError, err)
		return nil, err
	}
	return payload, nil
}
