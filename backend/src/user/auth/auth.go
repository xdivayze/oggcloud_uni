package auth

import (
	"time"

	"github.com/google/uuid"
)

const CODE_LENGTH = 32
const CODE_VALIDATION_LENGTH_MIN = 60
const AUTH_CODE_FIELDNAME = "authCode"

type AuthorizationCode struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Code string `gorm:"unique"`
	UserID uuid.UUID 
	ExpiresAt time.Time
}