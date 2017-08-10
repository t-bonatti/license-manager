package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/t-bonatti/license-manager/model"
)

func TestCreate(t *testing.T) {
	var assert = assert.New(t)
	var license = model.License{
		ID:      "abcdef",
		Version: "1",
		Info:    types.JSONText("{}"),
	}
	var ds = new(mockDatastore)
	ds.On("Create", mock.AnythingOfType("model.License")).Return(nil)
	var ts = httptest.NewServer(Create(ds))
	defer ts.Close()
	body, _ := json.Marshal(&license)
	res, err := http.Post(ts.URL, "application/json", bytes.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)
}

func TestCreateDsError(t *testing.T) {
	var assert = assert.New(t)
	var license = model.License{
		ID:      "abcdef",
		Version: "1",
		Info:    types.JSONText("{}"),
	}
	var ds = new(mockDatastore)
	ds.On("Create", mock.AnythingOfType("model.License")).Return(errors.New("error"))
	var ts = httptest.NewServer(Create(ds))
	defer ts.Close()
	body, _ := json.Marshal(&license)
	res, err := http.Post(ts.URL, "application/json", bytes.NewReader(body))
	assert.NoError(err)
	assert.Equal(http.StatusBadRequest, res.StatusCode)
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

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}/versions/{version}", Get(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	var respLicense model.License
	res, err := http.Get(ts.URL + "/abcdef/versions/1")
	assert.NoError(err)
	assert.NoError(json.NewDecoder(res.Body).Decode(&respLicense))
	assert.Equal(license, respLicense)
	assert.Equal(http.StatusOK, res.StatusCode)
}

func TestGetDsError(t *testing.T) {
	var assert = assert.New(t)
	var ds = new(mockDatastore)
	ds.On("Get", "abcdef", "1").Return(model.License{}, errors.New("fake err"))

	var mux = mux.NewRouter()
	mux.HandleFunc("/{id}/versions/{version}", Get(ds))
	var ts = httptest.NewServer(mux)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/abcdef/versions/1")
	assert.NoError(err)
	assert.Equal(http.StatusInternalServerError, res.StatusCode)
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
