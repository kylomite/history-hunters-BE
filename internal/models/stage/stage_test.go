package stage

import (
	"database/sql"
	"log"
	"testing"

	"historyHunters/internal/db"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sql.DB {
	err := godotenv.Load("../../../.env.test")
	if err != nil {
		log.Println("Failed to load .env file:", err)
	}

	testDB, err := db.ConnectDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	_, _ = testDB.Exec("DELETE FROM stages")
	_, _ = testDB.Exec("ALTER SEQUENCE stages_id_seq RESTART WITH 1")

	return testDB
}

func TestStageFields(t *testing.T) {
	stage := NewStage("Test Title", "background.png", 3)

	assert.Equal(t, "Test Title", stage.Title)
	assert.Equal(t, "background.png", stage.BackgroundImg)
	assert.Equal(t, 3, stage.Difficulty)
	assert.NotZero(t, stage.CreatedAt)
	assert.NotZero(t, stage.UpdatedAt)
}

func TestStageValidation(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage1 := NewStage("", "background.png", 5)
	err := stage1.Save(testDB)
	assert.EqualError(t, err, "title is required")

	stage2 := NewStage("Test Title", "", 5)
	err = stage2.Save(testDB)
	assert.EqualError(t, err, "background image is required")

	stage3 := NewStage("Test Title", "background.png", 0)
	err = stage3.Save(testDB)
	assert.EqualError(t, err, "difficulty is required")
}

func TestCreateStage(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage := NewStage("New Stage", "background.jpg", 2)
	err := stage.Save(testDB)

	assert.NoError(t, err)
	assert.NotZero(t, stage.ID)

	var count int
	err = testDB.QueryRow(`SELECT COUNT(*) FROM stages WHERE id = $1`, stage.ID).Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestGetAllStages(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage1 := NewStage("Stage 1", "bg1.png", 1)
	stage2 := NewStage("Stage 2", "bg2.png", 2)

	assert.NoError(t, stage1.Save(testDB))
	assert.NoError(t, stage2.Save(testDB))

	stages, err := GetAllStages(testDB)
	assert.NoError(t, err)
	assert.Len(t, stages, 2)

	assert.Equal(t, "Stage 1", stages[0].Title)
	assert.Equal(t, "bg1.png", stages[0].BackgroundImg)
	assert.Equal(t, 1, stages[0].Difficulty)

	assert.Equal(t, "Stage 2", stages[1].Title)
	assert.Equal(t, "bg2.png", stages[1].BackgroundImg)
	assert.Equal(t, 2, stages[1].Difficulty)
}

func TestFindStageByID(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage := NewStage("Stage Find", "bg_find.png", 4)
	assert.NoError(t, stage.Save(testDB))

	foundStage, err := FindStageByID(testDB, stage.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundStage)

	assert.Equal(t, stage.ID, foundStage.ID)
	assert.Equal(t, "Stage Find", foundStage.Title)
	assert.Equal(t, "bg_find.png", foundStage.BackgroundImg)
	assert.Equal(t, 4, foundStage.Difficulty)
}

func TestUpdateStage(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage := NewStage("Original Stage", "bg_orig.png", 2)
	assert.NoError(t, stage.Save(testDB))

	stage.Title = "Updated Stage"
	stage.BackgroundImg = "bg_updated.png"
	stage.Difficulty = 5

	err := stage.Update(testDB)
	assert.NoError(t, err)

	updatedStage, err := FindStageByID(testDB, stage.ID)
	assert.NoError(t, err)

	assert.Equal(t, "Updated Stage", updatedStage.Title)
	assert.Equal(t, "bg_updated.png", updatedStage.BackgroundImg)
	assert.Equal(t, 5, updatedStage.Difficulty)

}

func TestDeleteStage(t *testing.T) {
	testDB := setupTestDB(t)
	defer testDB.Close()

	stage := NewStage("Stage to Delete", "bg_delete.png", 3)
	assert.NoError(t, stage.Save(testDB))

	foundStage, err := FindStageByID(testDB, stage.ID)
	assert.NoError(t, err)
	assert.NotNil(t, foundStage)

	err = DeleteStage(testDB, stage.ID)
	assert.NoError(t, err)

	deletedStage, err := FindStageByID(testDB, stage.ID)
	assert.Nil(t, deletedStage)
	assert.Error(t, err)
	assert.Equal(t, "stage not found", err.Error())
}