-- name: CreateAppointment :one
INSERT INTO appointments (patient_id, user_id, visit_date, visit_timestamp, patient_notes, doctor_notes)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAppointmentByID :one
SELECT * FROM appointments WHERE id = $1;

-- name: GetAllAppointments :many
SELECT * FROM appointments
ORDER BY visit_date DESC;

-- name: GetAppointmentsByDate :many
SELECT * FROM appointments
WHERE visit_date = $1
ORDER BY appointment_sequence ASC;

-- name: GetAppointmentsByPatient :many
SELECT * FROM appointments
WHERE patient_id = $1
ORDER BY appointment_sequence ASC;

-- name: GetAppointmentBySequence :one
SELECT * FROM appointments
WHERE visit_date = $1 AND appointment_sequence = $2
ORDER BY created_at;

-- name: UpdateAppointment :one
UPDATE appointments
SET
    patient_notes = COALESCE($2, patient_notes),
    doctor_notes = COALESCE($3, doctor_notes),
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteAppointment :exec
DELETE FROM appointments WHERE id = $1;

