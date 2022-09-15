package database

import (
	"fmt"
	"mvc-testing/config"
	"mvc-testing/models"
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

func GetUserByID(id int64) (*models.UserResponse, error) {
	var user models.User

	result := db.First(&user, id)
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

	if count := tx.Where("email = ?", user.Email).First(&user).RowsAffected; count == 1 {
		tx.Rollback()
		return nil, fmt.Errorf("email already taken, please find another email")
	}

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

func DeleteUserByID(id int64) error {
	var user models.User

	tx := db.Begin()

	if _, err := GetUserByID(id); err != nil {
		tx.Rollback()
		return fmt.Errorf("user with the given ID not found")
	}

	if err := tx.Delete(&user, id).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func UpdateUserByID(userInput *models.UserInput, id int64) (*models.UserResponse, error) {
	var user models.User

	tx := db.Begin()

	if err := tx.First(&user, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := GetUserByID(id); err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("user with the given ID not found")
	}

	if count := tx.Where("email = ?", user.Email).Not("id = ?", id).RowsAffected; count == 1 {
		tx.Rollback()
		return nil, fmt.Errorf("email already taken, please find another email")
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
