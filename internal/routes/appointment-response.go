package routes

import (
	"patient-appointment-demo-go/internal/database"
	"time"
)

type AppointmentResponse struct {
	ID    int64 `json:"id"`
	PatientId    int64 `json:"patient_id"`
	UserId    int64 `json:"user_id"`
	VisitTime    time.Time `json:"visit_time"`
	VisitDate    time.Time `json:"visit_date"`
	PatientNotes string   `json:"patient_notes"`
	DoctorNotes  string `json:"doctor_notes"`
}

func AppointmentDbToResponse(data database.Appointment) AppointmentResponse{
    return AppointmentResponse {
        ID: int64(data.ID),
        PatientId: int64(data.PatientID),
        UserId: int64(data.UserID.Int32),
        VisitTime: data.VisitTimestamp.Time,
        VisitDate: data.VisitDate.Time,
        PatientNotes: data.PatientNotes.String,
        DoctorNotes: data.DoctorNotes.String,
    }
}

func AppointmentDbArrayToResponse(data []database.Appointment) []AppointmentResponse {

    appointments := make([]AppointmentResponse, len(data))

    for i,item := range data {
        appointments[i] = AppointmentDbToResponse(item)
    }

    return appointments

}
