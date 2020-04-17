package controllers

import (
	"net/http"
)

type Message struct {
	Message string
}

func BadRequestError(c Context, message string) {
	SendErrorResponse(c, http.StatusBadRequest, message)
}

func NotFoundError(c Context, message string) {
	SendErrorResponse(c, http.StatusNotFound, message)
}

func UnauthorizedError(c Context, message string) {
	SendErrorResponse(c, http.StatusUnauthorized, message)
}

func InternalServerError(c Context, message string) {
	SendErrorResponse(c, http.StatusInternalServerError, message)
}

func SendErrorResponse(c Context, status int, message string) {
	c.JSON(status, &Message{Message: message})
	c.Abort()
}
