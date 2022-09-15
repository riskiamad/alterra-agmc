package controllers

import (
	"mvc-middleware/lib/database"
	"mvc-middleware/middlewares"
	"mvc-middleware/models"
	"net/http"
	"strconv"

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

	if ok, err := middlewares.IsRequestValid(userInput); !ok {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "failed",
			"message": "body request not valid",
			"error":   middlewares.CustomValidationError(err),
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
	id, err := strconv.Atoi(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "the given id is not valid",
			"error":   err.Error(),
		})
	}

	result, err := database.GetUserByID(int64(id))
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
	id, err := strconv.Atoi(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "the given id is not valid",
			"error":   err.Error(),
		})
	}

	credentialID := middlewares.ExtractToken(ctx)

	if credentialID != int64(id) {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "you are not allowed to delete another user",
		})
	}

	err = database.DeleteUserByID(int64(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "user has deleted",
	})
}

func UpdateUserByID(ctx echo.Context) error {
	var err error

	userID := ctx.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "the given id is not valid",
			"error":   err.Error(),
		})
	}

	credentialID := middlewares.ExtractToken(ctx)

	if credentialID != int64(id) {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"status":  "failed",
			"message": "you are not allowed to edit another user",
		})
	}

	userInput := &models.UserInput{}
	err = ctx.Bind(&userInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to bind body request",
			"error":   err.Error(),
		})
	}

	if ok, err := middlewares.IsRequestValid(userInput); !ok {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "failed",
			"message": "body request not valid",
			"error":   middlewares.CustomValidationError(err),
		})
	}

	result, err := database.UpdateUserByID(userInput, int64(id))
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
