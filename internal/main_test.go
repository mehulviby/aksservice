package main

import (
	"aksservice/internal/api"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestAddAksConfigSuccess(t *testing.T) {
	controller := &api.Controller{}
	yfile, _ := ioutil.ReadFile("yamltest/valid1.yaml")
	req := httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	err := controller.AddAksConfig(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	assert.Equal(t, rec.Body.Len(), 0)
	assert.Equal(t, 201, rec.Code)

	yfile, _ = ioutil.ReadFile("yamltest/valid2.yaml")
	req = httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec = httptest.NewRecorder()
	err = controller.AddAksConfig(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	assert.Equal(t, rec.Body.Len(), 0)
	assert.Equal(t, 201, rec.Code)
}

func TestAddAksConfigError(t *testing.T) {
	controller := &api.Controller{}
	yfile, _ := ioutil.ReadFile("yamltest/invalid1.yaml")
	req := httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	err := controller.AddAksConfig(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	assert.Contains(t, rec.Body.String(), errors.New("Error:Field validation").Error())
	assert.NotEqual(t, 201, rec.Code)

	yfile, _ = ioutil.ReadFile("yamltest/invalid2.yaml")
	req = httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec = httptest.NewRecorder()
	err = controller.AddAksConfig(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	assert.Contains(t, rec.Body.String(), errors.New("Error:Field validation").Error())
	assert.NotEqual(t, 201, rec.Code)

	yfile, _ = ioutil.ReadFile("yamltest/invalid3.yaml")
	req = httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec = httptest.NewRecorder()
	err = controller.AddAksConfig(SetupEcho().NewContext(req, rec))
	assert.NotNil(t, err)
	assert.NotEqual(t, 201, rec.Code)
}

func TestGetAksConfigByTitleSuccess(t *testing.T) {
	controller := &api.Controller{}
	yfile, _ := ioutil.ReadFile("yamltest/valid1.yaml")
	req := httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	_ = controller.AddAksConfig(SetupEcho().NewContext(req, rec))

	req = httptest.NewRequest(http.MethodGet, "/aksconfig/Valid", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec = httptest.NewRecorder()
	err := controller.GetAksConfigByTitle(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	var result []interface{}
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, 1, len(result))

	yfile, _ = ioutil.ReadFile("yamltest/valid2.yaml")
	req = httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec = httptest.NewRecorder()
	err = controller.AddAksConfig(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	assert.Equal(t, rec.Body.Len(), 0)
	assert.Equal(t, 201, rec.Code)

	req = httptest.NewRequest(http.MethodGet, "/aksconfig?title=Valid", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec = httptest.NewRecorder()
	err = controller.GetAksConfigByTitle(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, 2, len(result))
}

func TestDeleteAllAksConfig(t *testing.T) {
	controller := &api.Controller{}
	yfile, _ := ioutil.ReadFile("yamltest/valid1.yaml")
	req := httptest.NewRequest(http.MethodPost, "/aksconfig", strings.NewReader(string(yfile)))
	req.Header.Set(echo.HeaderContentType, echo.MIMETextPlain)
	rec := httptest.NewRecorder()
	_ = controller.AddAksConfig(SetupEcho().NewContext(req, rec))

	req = httptest.NewRequest(http.MethodDelete, "/aksconfig", nil)
	rec = httptest.NewRecorder()
	_ = controller.DeleteAllAksConfig(SetupEcho().NewContext(req, rec))
	assert.Equal(t, 200, rec.Code)

	req = httptest.NewRequest(http.MethodGet, "/aksconfig", nil)
	rec = httptest.NewRecorder()
	err := controller.GetAllAksConfig(SetupEcho().NewContext(req, rec))
	assert.Nil(t, err)
	var result []interface{}
	_ = json.Unmarshal(rec.Body.Bytes(), &result)
	assert.Equal(t, 200, rec.Code)
	assert.Equal(t, 0, len(result))
}
