package model

import (
	"encoding/hex"
	"errors"
	"fmt"
	"oggcloudserver/src/db"
	"oggcloudserver/src/oggcrypto"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserFromMail(mail string) (*User, error) {
	user := User{}
	res := db.DB.Where("Email = ?", mail).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("record not found:\n\t%w", gorm.ErrRecordNotFound)
	} else if res.Error != nil {
		return nil, fmt.Errorf("error occured when querying the database:\n\t%w", res.Error)
	}
	return &user, nil
}

func GetUserFromID(id uuid.UUID) (*User, error) {
	user := User{}
	res := db.DB.Find(&user, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("record not found:\n\t%w", gorm.ErrRecordNotFound)
	} else if res.Error != nil {
		return nil, fmt.Errorf("error occured when querying the database:\n\t%w", res.Error)
	}
	return &user, nil
}

func GenerateAndEncryptSharedKey(clientpub string) (string, string, error) {

	pubkey, err := oggcrypto.ReadFromPEM(clientpub)
	if err != nil {
		return "", "", fmt.Errorf("error occured when reading from pem:\n\t%w", err)
	}

	ss, sp, err := oggcrypto.GenerateECDHPair()
	if err != nil {
		return "", "", fmt.Errorf("error when generating server ecdh pair:\n\t%w", err)
	}

	shared, salt, err := oggcrypto.DeriveSharedSecret(ss, pubkey, nil)
	if err != nil {
		return "", "", fmt.Errorf("error occured while deriving the shared secret:\n\t%w", err)
	}
	{
		sp, err := oggcrypto.EncodePublicKeyToPEM(sp)
		if err != nil {
			return "", "", fmt.Errorf("error occured while encoding server public key to pem:\n\t%w", err)
		}
		return hex.EncodeToString(salt) + hex.EncodeToString(shared), sp, nil
	}

}

func (u *User) ToString() string {
	return fmt.Sprintf("\tID:%s\n\tEmail:%s\n\tPasswordHash:%s\n\tEcdhSharedKey:%s\n\tCreatedAt:%s\n\tUpdatedAt:%s\n\t",
		u.ID.String(), u.Email, *u.PasswordHash, *u.EcdhSharedKey, u.CreatedAt.String(), u.UpdatedAt.String())

}
