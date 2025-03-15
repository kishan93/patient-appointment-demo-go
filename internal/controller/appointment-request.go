package controller

import "time"

type AppointmentCreateRequest struct {
	VisitTime    time.Time `json:"visit_time" validate:"required"`
	PatientNotes *string   `json:"patient_notes" validate:"omitempty,max=1000"`
}

type AppointmentUpdateRequest struct {
	PatientNotes *string `json:"patient_notes" validate:"omitempty,max=1000"`
	DoctorNotes  *string `json:"doctor_notes" validate:"omitempty,max=1000"`
}
