package controllers

import (
	"fmt"
	"log"
	"mvc-testing/middlewares"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	payloadUserBindingFailed = `{
		"fullname": "Rizky",
		"email": "rizky@gmail.com",
		"password": 123
	}`
	payloadUserSuccess = `{
		"fullname": "Rizky",
		"email": "riskiamad@gmail.com",
		"password": "Rahasia123"
	}`
	payloadUserEmailExist = `{
		"fullname": "Rizky",
		"email": "rizky@gmail.com",
		"password": "Rahasia123"
	}`
	payloadUserInvalid = `{
		"email": "rizky@gmail"
	}`
)

func TestGetAllUsersSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(GetAllUsers(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "success to fetch data from server")
		asserts.Contains(body, "success")
	}
}

func TestGetUserByIDNotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(GetUserByID(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to fetch data from server")
		asserts.Contains(body, "record not found")
	}
}

func TestGetUserByIDSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("4")

	asserts := assert.New(t)

	if asserts.NoError(GetUserByID(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "success")
	}
}

func TestGetUserByIDParamsNotValid(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("a")

	asserts := assert.New(t)

	if asserts.NoError(GetUserByID(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "the given id is not valid")
		asserts.Contains(body, "failed")
	}
}

func TestCreateUserBindingFailed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserBindingFailed))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewUser(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to bind body request")
		asserts.Contains(body, "failed")
	}
}

func TestCreateUserInvalidPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewUser(c)) {
		asserts.Equal(422, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{
			"error": {
				"UserInput.Email": "Email is not valid email",
				"UserInput.Fullname": "Fullname is required",
				"UserInput.Password": "Password is required"
			},
			"message": "body request not valid",
			"status": "failed"
		}`, body)
	}
}

func TestCreateUserEmailAlreadyExist(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserEmailExist))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewUser(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "email already taken, please find another email")
		asserts.Contains(body, "failed to fetch data from server")
	}
}

func TestCreateUserSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(CreateNewUser(c)) {
		asserts.Equal(201, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "user has created")
		asserts.Contains(body, "success")
	}
}

func TestUpdateUserByIDInvalidPayload(t *testing.T) {
	e := echo.New()
	token, err := middlewares.GenerateToken(4)
	if err != nil {
		log.Fatal(err)
	}
	println(token)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserInvalid))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	c.SetParamNames("id")
	c.SetParamValues("4")

	asserts := assert.New(t)

	if asserts.NoError(UpdateUserByID(c)) {
		asserts.Equal(422, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{
			"error": {
				"UserInput.Email": "Email is not valid email",
				"UserInput.Fullname": "Fullname is required",
				"UserInput.Password": "Password is required"
			},
			"message": "body request not valid",
			"status": "failed"
		}`, body)
	}
}

func TestUpdateUserByIDNotFound(t *testing.T) {
	e := echo.New()
	token, err := middlewares.GenerateToken(4)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Add("Authorization", "Bearer "+token)
	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)

	if asserts.NoError(UpdateUserByID(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "record not found")
	}
}

func TestUpdateUserByIDBindingFailed(t *testing.T) {
	e := echo.New()
	token, err := middlewares.GenerateToken(4)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserBindingFailed))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Add("Authorization", "Bearer "+token)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(UpdateUserByID(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to bind body request")
	}
}

func TestUpdateUserByIDSuccess(t *testing.T) {
	e := echo.New()
	token, err := middlewares.GenerateToken(4)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", strings.NewReader(payloadUserSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Add("Authorization", "Bearer "+token)
	c.SetParamNames("id")
	c.SetParamValues("4")

	asserts := assert.New(t)

	if asserts.NoError(UpdateUserByID(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "user has updated")
	}
}

func TestDeleteUserByIDNotFound(t *testing.T) {
	e := echo.New()
	token, err := middlewares.GenerateToken(4)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Add("Authorization", "Bearer "+token)
	c.SetParamNames("id")
	c.SetParamValues("2")

	asserts := assert.New(t)

	if asserts.NoError(DeleteUserByID(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "record not found")
	}
}

func TestDeleteUserByIDSuccess(t *testing.T) {
	e := echo.New()
	token, err := middlewares.GenerateToken(4)
	if err != nil {
		log.Fatal(err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/v1/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Request().Header.Add("Authorization", "Bearer "+token)
	c.SetParamNames("id")
	c.SetParamValues("1")

	asserts := assert.New(t)

	if asserts.NoError(DeleteUserByID(c)) {
		asserts.Equal(204, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "user has deleted")
	}
}
