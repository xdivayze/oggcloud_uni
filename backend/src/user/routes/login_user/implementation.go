package loginuser

import (
	"encoding/hex"
	"fmt"
	"oggcloudserver/src/user/model"

	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(passwordhex string, user *model.User) error {
	passwordbytes, err := hex.DecodeString(passwordhex)
	if err != nil {
		return fmt.Errorf("error occured while decoding password hex string:\n\t%w", err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), passwordbytes); err != nil {
		return fmt.Errorf("password no match, error occured:\n\t%w", err)
	}
	return nil

}
