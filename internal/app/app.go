package app

import (
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"metrics/internal/endpoint"
)

type App struct {
	e    *endpoint.EndPoint
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{
		echo: echo.New(),
		e:    endpoint.New(),
	}

	a.echo.Use(echoprometheus.NewMiddleware("app"))

	// a.echo.Use()

	a.echo.GET("/hello", a.e.Hello)
	a.echo.GET("/time", a.e.WhatTime)
	a.echo.GET("/metrics", echoprometheus.NewHandler())

	return a, nil
}

func (a *App) Run(port string) error {
	return a.echo.Start(port)
}
