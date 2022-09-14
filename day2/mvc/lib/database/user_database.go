package database

import (
	"fmt"
	"mvc/config"
	"mvc/models"
	"strconv"
)

var db = config.InitDB()

func GetAllUsers() (mx []*models.UserResponse, total int64, err error) {
	var users []models.User

	result := db.Find(&users)
	if err = result.Error; err != nil {
		return nil, 0, err
	}

	total = result.RowsAffected

	if total == 0 {
		return nil, total, fmt.Errorf("no record found")
	}

	for _, v := range users {
		userResponse := &models.UserResponse{
			ID:        v.ID,
			Fullname:  v.Fullname,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}

		mx = append(mx, userResponse)
	}

	return mx, total, nil
}

func GetUserByID(id string) (*models.UserResponse, error) {
	var user models.User

	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	result := db.First(&user, userId)
	if err := result.Error; err != nil {
		return nil, err
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user with the given ID not found")
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return &userResponse, nil
}

func CreateNewUser(userInput *models.UserInput) (*models.UserResponse, error) {

	user := models.User{
		Fullname: userInput.Fullname,
		Email:    userInput.Email,
		Password: userInput.Password,
	}

	tx := db.Begin()

	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	tx.Commit()

	return &userResponse, nil
}

func DeleteUserByID(id string) error {
	var user models.User

	userID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	tx := db.Begin()

	if count := tx.First(&user, userID).RowsAffected; count == 0 {
		tx.Rollback()
		return fmt.Errorf("user with the given ID not found")
	}

	if err := tx.Delete(&user, userID).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func UpdateUserByID(userInput *models.UserInput, id string) (*models.UserResponse, error) {
	var user models.User

	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	tx := db.Begin()

	if err := tx.First(&user, userID).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if count := tx.First(&user, userID).RowsAffected; count == 0 {
		tx.Rollback()
		return nil, fmt.Errorf("user with the given ID not found")
	}

	user.Fullname = userInput.Fullname
	user.Email = userInput.Email
	user.Password = userInput.Password

	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	userResponse := models.UserResponse{
		ID:        user.ID,
		Fullname:  user.Fullname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	tx.Commit()

	return &userResponse, nil
}
