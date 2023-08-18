package endpoint

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"metrics/internal/storage"
)

type EndPoint struct {
	s *storage.Storage
}

func New(s *storage.Storage) *EndPoint {
	return &EndPoint{
		s: s,
	}
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

func (e *EndPoint) AddUser(ctx echo.Context) error {
	name := ctx.Request().Header.Get("name")
	err := e.s.StoreUser(context.Background(), name)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.String(http.StatusOK, "User added!")
}

func (e *EndPoint) AllUsers(ctx echo.Context) error {
	names, err := e.s.AllUsers(context.Background())
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, names)
}
