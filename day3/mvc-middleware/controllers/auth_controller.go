package controllers

import (
	"mvc-middleware/lib/database"
	"mvc-middleware/middlewares"
	"mvc-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginUser(ctx echo.Context) error {
	var userLoginRequest models.UserLoginRequest
	var err error

	err = ctx.Bind(&userLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to bind body request",
			"error":   err.Error(),
		})
	}

	if ok, err := middlewares.IsRequestValid(userLoginRequest); !ok {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "failed",
			"message": "body request not valid",
			"error":   middlewares.CustomValidationError(err),
		})
	}

	result, err := database.LoginUser(&userLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "the credential match",
		"data":    result,
	})
}
