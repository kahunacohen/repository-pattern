-- name: GetEmployeeByLocalIdOrPassport :one
SELECT * from employees where local_id_number != '' AND LTRIM(local_id_number,'0') = ? OR foreign_passport_number != '' AND foreign_passport_number = ?;

-- name: UpdateEmployee :exec
UPDATE employees SET family_status_id=? WHERE id=?; 