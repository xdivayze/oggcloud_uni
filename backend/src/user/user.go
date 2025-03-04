package user

import (
	"encoding/hex"
	"fmt"
	"oggcloudserver/src/db"
	"oggcloudserver/src/user/auth/referral"
	"oggcloudserver/src/user/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//TODO collect all constants in a single file

var AdminID = "5fa510a6-dc4c-4a91-b726-08bc982104d7"
var AdminReferral string

func CreateAdminUser() error {

	passHash, err := bcrypt.GenerateFromPassword([]byte("hi man it's me"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error occurred while creating admin password hash:\n\t%w", err)
	}
	encodedPassHash := hex.EncodeToString(passHash)

	adminU := model.User{
		ID:           uuid.MustParse(AdminID),
		Email:        "admin@oggcloud.xyz",
		PasswordHash: &encodedPassHash,
	}

	if err := db.DB.Save(&adminU).Error; err != nil {
		return fmt.Errorf("error occurred while saving admin user:\n\t%w", err)
	}

	createdReferral, err := referral.CreateReferralModel()
	if err != nil {
		return fmt.Errorf("error occurred while creating admin referral")
	}
	AdminReferral = createdReferral.Code
	db.DB.Model(&adminU).Association("Referrals").Append(createdReferral)
	return nil
}
