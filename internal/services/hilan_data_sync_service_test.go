package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"testing"

	"github.com/kahunacohen/repo-pattern/internal/repositories"
)

func TestHilanDataSyncServiceSyncFlagNotSet(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()
	if err := loadSchemaAndSeed(db); err != nil {
		t.Fatalf("failed to load and seed db: %v", err)
	}

	// Unset flag in db.
	db.Exec("UPDATE companies SET employee_sync_active=false WHERE name='matav';")
	_, err := NewHilanDataSyncService(
		context.Background(),
		&repositories.CompanyImpl{DB: db},
		&repositories.EmployeeRepositoryImpl{DB: db},
		&repositories.FamilyStatusImpl{DB: db})
	if err.Error() != "error intializing HilanDataSyncService: company \"matav\" employee_sync_active flag not set" {
		t.Fatalf("did not get employee_sync_active not set error")
	}
}
func TestHilanDataSyncServiceSyncRecords(t *testing.T) {
	db, _ := sql.Open("sqlite3", ":memory:")
	defer db.Close()

	if err := loadSchemaAndSeed(db); err != nil {
		t.Fatalf("failed to load and seed db: %v", err)
	}
	// db.Exec("UPDATE companies SET employee_sync_active=true WHERE name='matav';")

	data, err := os.ReadFile("./MBTD594.json")
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	// Unmarshal a json file (representing hilan records parsed from a file)
	// so we can pass to SyncRecords.
	var records []hilanRecord
	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	hilanDataSyncService, err := NewHilanDataSyncService(
		context.Background(),
		&repositories.CompanyImpl{DB: db},
		&repositories.EmployeeRepositoryImpl{DB: db},
		&repositories.FamilyStatusImpl{DB: db})
	if err != nil {
		t.Fatal(err)
	}

	hilanDataSyncService.SyncRecords(records)
}
