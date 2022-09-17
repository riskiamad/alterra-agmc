package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	payloadBindingFailed = `{
		"Tittle": "Golang Advanced",
		"Author": "Alterra",
		"isbn": 3459797264
	}`
	payloadSuccess = `{"Tittle": "Golang Advanced",
	"Author": "Alterra",
	"isbn": "3459797264"
	}`
)

func TestGetAllBooksSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(GetAllBooks(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "success to fetch data from server")
		asserts.Contains(body, "success")
	}
}

func TestGetBookByIDNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(GetBookByID(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to fetch data from server")
		asserts.Contains(body, "book with the given id not found")
	}
}

func TestCreateBookBindingFailed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", strings.NewReader(payloadBindingFailed))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewBook(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to bind body request")
		asserts.Contains(body, "failed")
	}
}

func TestCreateBookInvalidPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewBook(c)) {
		asserts.Equal(422, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{
			"error": {
				"BookInput.Author": "Author is required",
				"BookInput.ISBN": "ISBN is required",
				"BookInput.Tittle": "Tittle is required"
			},
			"message": "body request not valid",
			"status": "failed"
		}`, body)
	}
}

func TestCreateBookSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", strings.NewReader(payloadSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewBook(c)) {
		asserts.Equal(201, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "book has created")
		asserts.Contains(body, "success")
	}
}

func TestGetBookByIDSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(GetBookByID(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "success")
	}
}

func TestUpdateBookByIDInvalidPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(UpdateBookByID(c)) {
		asserts.Equal(422, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{
			"error": {
				"BookInput.Author": "Author is required",
				"BookInput.ISBN": "ISBN is required",
				"BookInput.Tittle": "Tittle is required"
			},
			"message": "body request not valid",
			"status": "failed"
		}`, body)
	}
}

func TestUpdateBookByIDNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", strings.NewReader(payloadSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)

	if asserts.NoError(UpdateBookByID(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "book with the given id not found")
	}
}

func TestUpdateBookByIDBindingFailed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", strings.NewReader(payloadBindingFailed))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(UpdateBookByID(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to bind body request")
	}
}

func TestUpdateBookByIDSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", strings.NewReader(payloadSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(UpdateBookByID(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "book has updated")
	}
}

func TestDeleteBookByIDNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)

	if asserts.NoError(DeleteBookByID(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "book with the given id not found")
	}
}

func TestDeleteBookByIDSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/books", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(DeleteBookByID(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "book has deleted")
	}
}
