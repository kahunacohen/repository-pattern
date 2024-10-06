-- name: GetPatient :one
SELECT * from patients where id = ?;

-- name: GetPatients :many
SELECT * FROM patients;

-- name: GetPatientsWithEmergencyDetails :many
SELECT * FROM patients_with_emergency_details;