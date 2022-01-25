package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/t-bonatti/license-manager/model"
)

func TestCreate(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockDatastore)
	ds.On("Create", mock.AnythingOfType("model.License")).Return(nil)
	rPath := "/license"
	r := gin.Default()
	r.POST(rPath, Create(ds))
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{"id": "abcdef","version": "1","info":{}}`))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(http.StatusCreated, w.Code)
}

func TestCreateDsError(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockDatastore)
	ds.On("Create", mock.AnythingOfType("model.License")).Return(errors.New("error"))
	rPath := "/license"
	r := gin.Default()
	r.POST(rPath, Create(ds))
	req, _ := http.NewRequest("POST", rPath, strings.NewReader(`{"id": "abcdef","version": "1","info":{}}`))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(http.StatusBadRequest, w.Code)
}

func TestGet(t *testing.T) {
	var assert = assert.New(t)
	var license = model.License{
		ID:      "abcdef",
		Version: "1",
		Info:    types.JSONText("{}"),
	}
	var ds = new(mockDatastore)
	ds.On("Get", "abcdef", "1").Return(license, nil)

	r := gin.Default()
	r.GET("/license/:id/versions/:version", Get(ds))
	req, _ := http.NewRequest("GET", "/license/abcdef/versions/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var respLicense model.License
	assert.NoError(json.NewDecoder(w.Body).Decode(&respLicense))
	assert.Equal(license, respLicense)
	assert.Equal(http.StatusOK, w.Code)
}

func TestGetDsError(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockDatastore)
	ds.On("Get", "abcdef", "1").Return(model.License{}, errors.New("fake err"))

	r := gin.Default()
	r.GET("/license/:id/versions/:version", Get(ds))
	req, _ := http.NewRequest("GET", "/license/abcdef/versions/1", strings.NewReader(`{"id": "1","name": "joe"}`))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(http.StatusInternalServerError, w.Code)
}

type mockDatastore struct {
	mock.Mock
}

func (m mockDatastore) Create(license model.License) (err error) {
	var args = m.Called(license)
	return args.Error(0)
}

func (m mockDatastore) Get(id string, version string) (license model.License, err error) {
	var args = m.Called(id, version)
	return args.Get(0).(model.License), args.Error(1)
}
