package model

import "github.com/google/uuid"

type Role struct {
	BaseModel
	ID   uuid.UUID `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name string    `json:"name"`
}
