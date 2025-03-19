package user

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"oggcloudserver/src/db"
	"oggcloudserver/src/oggcrypto"
	"oggcloudserver/src/user/auth/referral"
	"oggcloudserver/src/user/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//TODO collect all constants in a single file

var AdminID = "5fa510a6-dc4c-4a91-b726-08bc982104d7"

const ADMIN_PASSWORD = "osiman34"

var AdminReferral string

func CreateAdminUser() error {

	passHash, err := oggcrypto.CalculateSHA256sum(bytes.NewReader([]byte(ADMIN_PASSWORD)))
	if err != nil {
		return fmt.Errorf("error occurred while generating pass hash:\n\t%w ", err)
	}
	passHashBytes,err := hex.DecodeString(passHash)
	if  err != nil {
		return fmt.Errorf("error occurred while decoding hexadecimal string:\n\t%w", err)
	}
	passBcrypt, err := bcrypt.GenerateFromPassword(passHashBytes, bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("error occurred while creating admin password hash:\n\t%w", err)
	}
	encodedPassHash := string(passBcrypt)

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
