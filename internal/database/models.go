// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Appointment struct {
	ID                  int32
	PatientID           int32
	UserID              pgtype.Int4
	VisitDate           pgtype.Date
	AppointmentSequence int16
	VisitTimestamp      pgtype.Timestamptz
	PatientNotes        pgtype.Text
	DoctorNotes         pgtype.Text
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
}

type Patient struct {
	ID        int32
	Name      string
	Phone     pgtype.Text
	Email     string
	Age       pgtype.Int2
	Weight    pgtype.Numeric
	Height    pgtype.Numeric
	Gender    pgtype.Text
	Address   pgtype.Text
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}

type User struct {
	ID        int32
	Email     string
	Password  string
	Type      string
	CreatedAt pgtype.Timestamp
	UpdatedAt pgtype.Timestamp
}
