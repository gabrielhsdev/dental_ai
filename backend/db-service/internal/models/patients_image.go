package models

import (
	"github.com/google/uuid"
)

type PatientImages struct {
	Id          uuid.UUID `json:"id"`
	PatientId   uuid.UUID `json:"patientId"`
	ImageData   string    `json:"imageData"` // Path to the image file in the storage
	FileType    string    `json:"fileType"`  // 3 Possible values: "jpg", "png" or "png"
	Description string    `json:"description"`
	UploadedAt  string    `json:"uploadedAt"`
	CreatedAt   string    `json:"createdAt"`
	UpdatedAt   string    `json:"updatedAt"`
}
