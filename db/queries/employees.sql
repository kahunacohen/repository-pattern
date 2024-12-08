-- name: GetByLocalIdOrPassport :one
SELECT * from employees where local_id_number != '' AND LTRIM(local_id_number,'0') = ? OR foreign_passport_number != '' AND foreign_passport_number = ?;