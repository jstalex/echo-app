package app

import (
	"log"

	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"metrics/internal/endpoint"
	"metrics/internal/metrics"
	"metrics/internal/middleware"
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

	a.echo.Use(echoprometheus.NewMiddleware("metric_app"))

	err := prometheus.Register(metrics.NameMetric)
	if err != nil {
		log.Printf(err.Error())
	}

	a.echo.GET("/hello", a.e.Hello)
	a.echo.GET("/time", a.e.WhatTime)
	a.echo.GET("/metrics", echoprometheus.NewHandler())
	a.echo.POST("/name", a.e.Name, middleware.UserCheck)

	return a, nil
}

func (a *App) Run(port string) error {
	return a.echo.Start(port)
}
