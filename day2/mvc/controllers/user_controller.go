package controllers

import (
	"mvc/lib/database"
	"mvc/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(ctx echo.Context) error {
	result, total, err := database.GetAllUsers()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success to fetch data from server",
		"data":    result,
		"total":   total,
	})
}

func CreateNewUser(ctx echo.Context) error {
	var err error

	userInput := &models.UserInput{}
	err = ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to bind body request",
			"error":   err.Error(),
		})
	}

	result, err := database.CreateNewUser(userInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "user has created",
		"data":    result,
	})
}

func GetUserByID(ctx echo.Context) error {
	userID := ctx.Param("id")
	result, err := database.GetUserByID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "success to fetch data from server",
		"data":    result,
	})
}

func DeleteUserByID(ctx echo.Context) error {
	userID := ctx.Param("id")
	err := database.DeleteUserByID(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "book has deleted",
	})
}

func UpdateUserByID(ctx echo.Context) error {
	var err error

	userID := ctx.Param("id")
	userInput := &models.UserInput{}
	err = ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to bind body request",
			"error":   err.Error(),
		})
	}

	result, err := database.UpdateUserByID(userInput, userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user has updated",
		"data":    result,
	})
}
