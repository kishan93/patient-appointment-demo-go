-- +goose Up
CREATE TABLE appointments (
    id SERIAL PRIMARY KEY,
    patient_id INT NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    user_id INT REFERENCES users(id) ON DELETE SET NULL,
    visit_date DATE NOT NULL,
    appointment_sequence SMALLINT NOT NULL,
    visit_timestamp TIMESTAMPTZ NOT NULL,
    patient_notes TEXT,
    doctor_notes TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    UNIQUE (visit_date, appointment_sequence) -- Ensures unique sequence per day
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION generate_appointment_sequence()
RETURNS TRIGGER AS $$
DECLARE
    max_sequence SMALLINT;
BEGIN
    SELECT COALESCE(MAX(appointment_sequence), 0) + 1 INTO max_sequence
    FROM appointments
    WHERE visit_date = NEW.visit_date;

    NEW.appointment_sequence := max_sequence;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +goose StatementEnd


CREATE TRIGGER update_updated_at_on_appointments_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();


CREATE TRIGGER set_appointment_sequence
BEFORE INSERT ON appointments
FOR EACH ROW
EXECUTE FUNCTION generate_appointment_sequence();

-- +goose Down
DROP TRIGGER IF EXISTS update_updated_at_on_appointments_trigger ON appointments;
DROP TRIGGER IF EXISTS set_appointment_sequence ON appointments;
DROP FUNCTION IF EXISTS generate_appointment_sequence;
DROP TABLE IF EXISTS appointments;

