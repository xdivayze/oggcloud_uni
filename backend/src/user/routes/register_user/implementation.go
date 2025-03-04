package registeruser

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"oggcloudserver/src/db"
	"oggcloudserver/src/user/auth/referral/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const PASSWORD_LENGTH = 9

var ErrPasswordTooLong = fmt.Errorf("password length exceeds %d characters", PASSWORD_LENGTH)

func processPassword(c *gin.Context, passwordhex string) (string, error) {
	hexpass, err := hex.DecodeString(passwordhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured registering user"})
		return "", fmt.Errorf("error occured while generating bytes from hex:\n\t%w", err)
	}

	if len(hexpass) > PASSWORD_LENGTH*8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrPasswordTooLong.Error()})
		return "", ErrPasswordTooLong
	}

	bcryptPass, err := bcrypt.GenerateFromPassword(hexpass, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured registering user"})
		return "", fmt.Errorf("error occured while generating bcrypt:\n\t%w", err)
	}
	return string(bcryptPass), nil

}



func processReferral(referralCode string,  id uuid.UUID,c *gin.Context) (bool, error) {
	var supposedReferral = model.Referral{}
	if err := db.DB.Where("code = ?", referralCode).Where("used = ?", false).First(&supposedReferral).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusForbidden)
			return false, nil
		}
		c.Status(http.StatusInternalServerError)
		return false,fmt.Errorf("err occurred while getting instance from db:\n\t%w", err)
	}
	supposedReferral.Used = true
	supposedReferral.AcceptedBy = &id
	if err := db.DB.Save(&supposedReferral).Error; err != nil {
		c.Status(http.StatusInternalServerError)
		return false, fmt.Errorf("error occurred while saving instance to db:\n\t%w ", err)
	}
	return true, nil
} 
