package routes

type PatientCreateRequest struct {
	Name    string  `json:"name" validate:"required,max=255"`
	Phone   string  `json:"phone" validate:"required,max=20"`
	Email   string  `json:"email" validate:"required,email,max=255"`
	Age     int16   `json:"age" validate:"omitempty,min=0"`
	Weight  float64 `json:"weight" validate:"omitempty,gt=0"`
	Height  float64 `json:"height" validate:"omitempty,gt=0"`
	Gender  string  `json:"gender" validate:"required,oneof=Male Female Other"`
	Address string  `json:"address" validate:"omitempty,max=500"`
}

type PatientUpdateRequest struct {
	Name    string  `json:"name" validate:"omitempty,max=255"`
	Phone   string  `json:"phone" validate:"omitempty,max=20"`
	Email   string  `json:"email" validate:"omitempty,email,max=255"`
	Age     int16   `json:"age" validate:"omitempty,min=0"`
	Weight  float64 `json:"weight" validate:"omitempty"`
	Height  float64 `json:"height" validate:"omitempty"`
	Gender  string  `json:"gender" validate:"omitempty,oneof=Male Female Other"`
	Address string  `json:"address" validate:"omitempty,max=500"`
}

