package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`

	Name string `json:"name"`
	Description string `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Members []User    `json:"members" gorm:"many2many:project_members;"`

	OwnerID uuid.UUID `json:"owner_id"`
	Owner User `json:"owner" gorm:"foreignKey:OwnerID"`
}

type ProjectMember struct {
	ProjectID uuid.UUID `gorm:"type:uuid"`
	UserID    uuid.UUID `gorm:"type:uuid"`
}