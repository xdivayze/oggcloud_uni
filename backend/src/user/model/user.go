package model

import (
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/file_ops/session/Services/upload"
	"oggcloudserver/src/user/auth"
	referral_model "oggcloudserver/src/user/auth/referral/model"
	"time"

	"github.com/google/uuid"
)




type User struct {
	ID                 uuid.UUID `gorm:"type:uuid;primaryKey"`
	Email              string    `gorm:"unique"`
	PasswordHash       *string
	EcdhSharedKey      *string
	AuthorizationCodes []auth.AuthorizationCode  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Sessions           []upload.Session          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Files              []file.File               `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Referrals          []referral_model.Referral `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:CreatedBy"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
}
