package model

import (
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/file_ops/session/Services/upload"
	"oggcloudserver/src/user/auth"
	"time"

	"github.com/google/uuid"
)

const PASSWORD_FIELDNAME = "password"
const EMAIL_FIELDNAME = "email"
const ECDH_PUB_FIELDNAME = "ecdh_public"

type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email              string    `gorm:"unique"`
	PasswordHash       *string
	EcdhSharedKey      *string
	AuthorizationCodes []auth.AuthorizationCode `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Sessions           []upload.Session         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Files              []file.File              `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
