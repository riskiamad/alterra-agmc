package controllers

import (
	"mvc-middleware/lib/database"
	"mvc-middleware/middlewares"
	"mvc-middleware/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllBooks(ctx echo.Context) error {
	result, total, err := database.GetAllBooks()
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

func CreateNewBook(ctx echo.Context) error {
	var err error

	bookInput := &models.BookInput{}
	err = ctx.Bind(&bookInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to bind body request",
			"error":   err.Error(),
		})
	}

	if ok, err := middlewares.IsRequestValid(bookInput); !ok {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "failed",
			"message": "body request not valid",
			"error":   middlewares.CustomValidationError(err),
		})
	}

	result, err := database.CreateNewBook(bookInput)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "book has created",
		"data":    result,
	})
}

func GetBookByID(ctx echo.Context) error {
	bookID := ctx.Param("id")
	result, err := database.GetBookByID(bookID)
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

func DeleteBookByID(ctx echo.Context) error {
	bookID := ctx.Param("id")
	result, err := database.DeleteBookByID(bookID)
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
		"data":    result,
	})
}

func UpdateBookByID(ctx echo.Context) error {
	var err error

	bookID := ctx.Param("id")
	bookInput := &models.BookInput{}
	err = ctx.Bind(&bookInput)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "failed to bind body request",
			"error":   err.Error(),
		})
	}

	if ok, err := middlewares.IsRequestValid(bookInput); !ok {
		return ctx.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"status":  "failed",
			"message": "body request not valid",
			"error":   middlewares.CustomValidationError(err),
		})
	}

	result, err := database.UpdateBookByID(bookInput, bookID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"status":  "failed",
			"message": "failed to fetch data from server",
			"error":   err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "book has updated",
		"data":    result,
	})
}
