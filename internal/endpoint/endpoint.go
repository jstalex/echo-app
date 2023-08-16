package endpoint

import (
	"net/http"
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
