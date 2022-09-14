package database

import (
	"fmt"
	"mvc/models"
	"strconv"
	"time"
)

var data = []models.Book{}

func GetAllBooks() ([]models.Book, int, error) {
	total := len(data)
	return data, total, nil
}

func CreateNewBook(bookInput *models.BookInput) (*models.Book, error) {

	var generateNewID int64
	if len(data) > 0 {
		generateNewID = data[len(data)-1].ID + 1
	} else {
		generateNewID = 1
	}

	book := models.Book{
		ID:        generateNewID,
		Tittle:    bookInput.Tittle,
		Author:    bookInput.Author,
		ISBN:      bookInput.ISBN,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	data = append(data, book)
	return &book, nil
}

func GetBookByID(id string) (m *models.Book, err error) {
	var count int
	var book models.Book
	bookID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		if v.ID == int64(bookID) {
			book = v
			count++
			fmt.Println(v)
		}
	}

	if count < 1 {
		return nil, fmt.Errorf("book with the given id not found")
	}

	return &book, nil
}

func DeleteBookByID(id string) (book *models.Book, err error) {
	var count, key int
	bookID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for k, v := range data {
		if v.ID == int64(bookID) {
			book = &v
			count++
			key = k
		}
	}

	if count < 1 {
		return nil, fmt.Errorf("book with the given id not found")
	}

	data = append(data[:key], data[key+1:]...)

	return book, nil
}

func UpdateBookByID(bookInput *models.BookInput, id string) (book *models.Book, err error) {
	var count int

	bookID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	for _, v := range data {
		if v.ID == int64(bookID) {
			count++
			v.Author = bookInput.Author
			v.Tittle = bookInput.Tittle
			v.ISBN = bookInput.ISBN
			v.UpdatedAt = time.Now()
			book = &v
		}
	}

	if count < 1 {
		return nil, fmt.Errorf("book with the given id not found")
	}

	return book, nil
}
