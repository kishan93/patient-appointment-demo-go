-- +goose Up
CREATE TABLE IF NOT EXISTS public.users
(
    id SERIAL PRIMARY KEY,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    password character(255) COLLATE pg_catalog."default" NOT NULL,
    type character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    updated_at timestamp without time zone,
    CONSTRAINT users_email_key UNIQUE (email)
);

CREATE TRIGGER update_updated_at_on_users_trigger
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();


-- +goose Down
DROP TRIGGER IF EXISTS update_updated_at_on_users_trigger ON users;
DROP TABLE IF EXISTS public.users;

