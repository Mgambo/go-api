package internal_models

import (
	"time"

	"github.com/google/uuid"
)

type DatabaseBaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey;default:uuid_generate_v4();"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at", gorm:"nullable:true;default:null"`
}
