-- name: CreatePatient :one
INSERT INTO patients (name, phone, email, age, weight, height, gender, address)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPatientByID :one
SELECT * FROM patients WHERE id = $1;

-- name: GetAllPatients :many
SELECT * FROM patients
WHERE (@name::text = '' OR name::text ILIKE '%' || @name::text || '%')
ORDER BY
    CASE
        WHEN @sort_by::text = 'name' THEN name
        WHEN @sort_by::text = 'age' THEN age::TEXT
        ELSE created_at::TEXT
    END
    || CASE WHEN @sort_direction::text = 'DESC' THEN ' DESC' ELSE ' ASC' END;

-- name: UpdatePatient :one
UPDATE patients
SET
    name = COALESCE($2, name),
    phone = COALESCE($3, phone),
    email = COALESCE($4, email),
    age = COALESCE($5, age),
    weight = COALESCE($6, weight),
    height = COALESCE($7, height),
    gender = COALESCE($8, gender),
    address = COALESCE($9, address),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePatient :exec
DELETE FROM patients WHERE id = $1;

