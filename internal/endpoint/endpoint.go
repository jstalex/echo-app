package endpoint

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

type EndPoint struct {
}

func New() *EndPoint {
	return &EndPoint{}
}

func (e *EndPoint) Hello(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello!")
}

func (e *EndPoint) WhatTime(ctx echo.Context) error {
	return ctx.String(http.StatusOK, time.Now().String())
}

func (e *EndPoint) Name(ctx echo.Context) error {
	name := ctx.Request().Header.Get("name")
	if strings.EqualFold(name, "vasya") {
		return ctx.String(http.StatusOK, "Welcome!")
	}
	return ctx.String(http.StatusBadRequest, "You are not Vasya")
}
