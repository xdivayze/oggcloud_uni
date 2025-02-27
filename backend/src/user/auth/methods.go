package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"oggcloudserver/src/db"
	"time"

	"github.com/google/uuid"
)

func CreateInstance(userID uuid.UUID) (*AuthorizationCode, error) {
	bytes := make([]byte, CODE_LENGTH)
	if _, err := rand.Read(bytes); err != nil {
		return nil, fmt.Errorf("error occured while reading from random buffer:\n\t%w", err)
	}
	validUntil := time.Now().Add(CODE_VALIDATION_LENGTH_MIN * time.Minute)
	return &AuthorizationCode{
		ID:        uuid.New(),
		UserID:    userID,
		Code:      hex.EncodeToString(bytes),
		ExpiresAt: validUntil,
	}, nil
}

func (a *AuthorizationCode) IsValid(destroy bool)(bool) {
	diff := time.Until(a.ExpiresAt).Milliseconds()
	if diff <= 0 {
		if destroy {
			db.DB.Delete(a)
		}
		return false
	}
	return true
}

func CheckValidity(code string)(bool, error){
	auth, err := RetrieveFromDB(code)
	if err != nil {
		return false, err
	}
	diff := time.Until(auth.ExpiresAt).Milliseconds()
	if diff <= 0 {
		return false, nil
	}
	return true, nil

}

func RetrieveFromDB(code string) (*AuthorizationCode,error) {
	auth := AuthorizationCode{}
	res := db.DB.Where("Code = ?", code).First(&auth)
	if res.Error != nil {
		return nil, fmt.Errorf("error occured while retrieving instance from db:\n\t%w", res.Error)
	}
	return &auth, nil
}

func DeleteFromDB(inst *AuthorizationCode) error {
	res := db.DB.Delete(inst)
	return res.Error
}

func SaveToDB(inst *AuthorizationCode) (error) {
	db := db.DB
	if res := db.Create(inst); res.Error != nil {
		return fmt.Errorf("error occured when trying to insert value:\n\t%w", res.Error)
	}
	return nil
	
}
