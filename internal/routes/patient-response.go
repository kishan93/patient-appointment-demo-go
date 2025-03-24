package routes

import "patient-appointment-demo-go/internal/database"

type PatientResponse struct {
	ID    int64  `json:"id"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	Email   string  `json:"email"`
	Age     int16   `json:"age"`
	Weight  float64 `json:"weight"`
	Height  float64 `json:"height"`
	Gender  string  `json:"gender"`
	Address string  `json:"address"`
}

func PatientDbToResponse(data database.Patient) PatientResponse {
	weight, _ := data.Weight.Float64Value()
	height, _ := data.Height.Float64Value()

	return PatientResponse{
		ID:   int64(data.ID),
		Name:   data.Name,
		Phone:  data.Phone.String,
		Email:  data.Email,
		Age:    data.Age.Int16,
		Weight: weight.Float64,
		Height: height.Float64,
        Gender: data.Gender.String,
        Address: data.Address.String,
	}
}

func PatientDbArrayToResponse(data []database.Patient) []PatientResponse {

    patients := make([]PatientResponse, len(data))

    for i,item := range data {
        patients[i] = PatientDbToResponse(item)
    }

    return patients

}
