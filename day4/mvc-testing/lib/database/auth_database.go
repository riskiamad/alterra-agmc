package database

import (
	"fmt"
	"mvc-testing/middlewares"
	"mvc-testing/models"
)

func LoginUser(userLoginRequest *models.UserLoginRequest) (*models.UserResponse, error) {
	var user models.User
	result := db.Where("email = ?", userLoginRequest.Email).Where("password = ?", userLoginRequest.Password).First(&user)

	if err := result.Error; err != nil {
		return nil, err
	}

	if count := result.RowsAffected; count == 0 {
		return nil, fmt.Errorf("credential not match")
	}

	token, err := middlewares.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userResponse, nil
}
