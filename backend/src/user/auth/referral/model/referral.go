package model

import (
	"time"

	"github.com/google/uuid"
)

type Referral struct {
	ID         uuid.UUID
	Code string
	CreatedBy  uuid.UUID
	AcceptedBy *uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
