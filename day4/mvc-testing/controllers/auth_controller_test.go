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
	payloadInvalidJSON     = `{"email":"test.com"}`
	matchCredentialPayload = `{
		"email":"rizky@gmail.com",
		"password":"Rahasia123"
	}`
	unmatchCredentialPayload = `{
		"email":"riski@gmail.com",
		"password":"Rahasia123"
	}`
	failedToBindPayload = `{
		"email":1,
		"password":"Rahasia123"
	}`
)

func TestLoginUserInvalidPayload(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(payloadInvalidJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(LoginUser(c)) {
		asserts.Equal(422, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{
			"error": {
				"UserLoginRequest.Email": "Email is not valid email",
				"UserLoginRequest.Password": "Password is required"
			},
			"message": "body request not valid",
			"status": "failed"
		}`, body)
	}
}

func TestLoginUserUnmatchCredential(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(unmatchCredentialPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(LoginUser(c)) {
		asserts.Equal(500, rec.Code)
		body := rec.Body.String()
		asserts.JSONEq(`{
			"error": "record not found",
			"message": "failed to fetch data from server",
			"status": "failed"
		}`, body)
	}
}

func TestLoginUserBindingFailed(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(failedToBindPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(LoginUser(c)) {
		asserts.Equal(400, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "failed to bind body request")
		asserts.Contains(body, "failed")
	}
}

func TestLoginUserSuccess(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(matchCredentialPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	asserts := assert.New(t)

	if asserts.NoError(LoginUser(c)) {
		asserts.Equal(200, rec.Code)
		body := rec.Body.String()
		asserts.Contains(body, "the credential match")
		asserts.Contains(body, "success")
		asserts.Contains(body, "token")
	}
}
