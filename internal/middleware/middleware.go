package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"
	"metrics/internal/metrics"
)

func UserCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		name := ctx.Request().Header.Get("name")
		if strings.EqualFold(name, "vasya") {
			metrics.AddVasya()
		}

		return next(ctx)
	}
}
