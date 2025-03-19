package login_user

import (
	"encoding/hex"
	"fmt"
	"oggcloudserver/src/user/model"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(passwordHex string, user *model.User) error {
	passwordBytes, err := hex.DecodeString(passwordHex)
	if err != nil {
		return fmt.Errorf("error occurred while decoding password hex string:\n\t%w", err)
	}
	bcryptHash := []byte(*user.PasswordHash)
	if err = bcrypt.CompareHashAndPassword(bcryptHash, passwordBytes); err != nil {
		return fmt.Errorf("password no match, error occurred:\n\t%w", err)
	}
	return nil

}
