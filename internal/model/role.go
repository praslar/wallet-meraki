package model

import "github.com/google/uuid"

type Role struct {
	BaseModel
	RoleID uuid.UUID `json:"role_id" gorm:"primaryKey"`
	Name   string    `json:"name"`
	Value  int       `json:"value"`
	Key    string    `json:"key"`
}
