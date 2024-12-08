package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

func TestHilanDataSyncService(t *testing.T) {
	data, err := os.ReadFile("./MBTD594.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	var records []hilanRecord

	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()
	employeeRepo := &repositories.EmployeeRepositoryImpl{DB: db}

	hilanDataSyncService := NewHilanDataSyncService(employeeRepo)
	hilanDataSyncService.SyncRecords(records)
}
