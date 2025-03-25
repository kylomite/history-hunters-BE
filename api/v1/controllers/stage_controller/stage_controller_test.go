package stage_controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"historyHunters/internal/db"
	"historyHunters/internal/models/stage"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/go-chi/chi/v5"
	"context"
)

func setupTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../../../.env.test")
	assert.NoError(t, err)

	testDB, err := db.ConnectDB()
	assert.NoError(t, err)

	_, err = testDB.Exec("TRUNCATE stages RESTART IDENTITY CASCADE")
	assert.NoError(t, err)

	return testDB
}

func TestCreateStage(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	handler := CreateStage(testDB)

	reqBody := `{"title": "Test Stage", "background_img": "bg1.png", "difficulty": 3}`
	req := httptest.NewRequest("POST", "/stages", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var createdStage stage.Stage
	err := json.NewDecoder(resp.Body).Decode(&createdStage)
	assert.NoError(t, err)

	assert.Equal(t, "Test Stage", createdStage.Title)
	assert.Equal(t, "bg1.png", createdStage.BackgroundImg)
	assert.Equal(t, 3, createdStage.Difficulty)
}

func TestGetAllStages(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage1 := stage.NewStage("Stage 1", "bg1.png", 2)
	assert.NoError(t, stage1.Save(testDB))

	stage2 := stage.NewStage("Stage 2", "bg2.png", 4)
	assert.NoError(t, stage2.Save(testDB))

	handler := GetAllStages(testDB)

	req := httptest.NewRequest("GET", "/stages", nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var stages []stage.Stage
	err := json.NewDecoder(resp.Body).Decode(&stages)
	assert.NoError(t, err)

	assert.Len(t, stages, 2)
	assert.Equal(t, "Stage 1", stages[0].Title)
	assert.Equal(t, "Stage 2", stages[1].Title)
}

func TestGetStageByID(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	testStage := stage.NewStage("Stage 1", "bg1.png", 2)
	assert.NoError(t, testStage.Save(testDB))

	router := chi.NewRouter()
	router.Get("/stages/{id}", GetStageByID(testDB))

	req := httptest.NewRequest("GET", "/stages/"+strconv.Itoa(testStage.ID), nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))
	chi.URLParamFromCtx(req.Context(), "id")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var fetchedStage stage.Stage
	err := json.NewDecoder(resp.Body).Decode(&fetchedStage)
	assert.NoError(t, err)

	assert.Equal(t, "Stage 1", fetchedStage.Title)
	assert.Equal(t, "bg1.png", fetchedStage.BackgroundImg)
	assert.Equal(t, 2, fetchedStage.Difficulty)
}

func TestDeleteStage(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	testStage := stage.NewStage("Stage to Delete", "bg_delete.png", 1)
	assert.NoError(t, testStage.Save(testDB))

	router := chi.NewRouter()
	router.Delete("/stages/{id}", DeleteStage(testDB))

	req := httptest.NewRequest("DELETE", "/stages/"+strconv.Itoa(testStage.ID), nil)
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chi.NewRouteContext()))
	chi.URLParamFromCtx(req.Context(), "id")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNoContent, resp.Code)

	_, err := stage.FindStageByID(testDB, testStage.ID)
	assert.Error(t, err)
}