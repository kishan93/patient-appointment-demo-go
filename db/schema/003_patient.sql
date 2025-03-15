-- +goose Up
CREATE TABLE IF NOT EXISTS public.patients
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(20),
    email VARCHAR(255) UNIQUE NOT NULL,
    age SMALLINT CHECK (age >= 0),
    weight DECIMAL(5,2),
    height DECIMAL(5,2),
    gender VARCHAR(10),
    address TEXT,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP,
    CONSTRAINT patients_email_key UNIQUE (email)
);

CREATE TRIGGER update_updated_at_on_patients_trigger
BEFORE UPDATE ON patients
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();


-- +goose Down
DROP TRIGGER IF EXISTS update_updated_at_on_patients_trigger ON patients;
DROP TABLE IF EXISTS public.patients;

