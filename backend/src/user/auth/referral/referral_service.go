package referral

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"oggcloudserver/src/db"
	ref_model "oggcloudserver/src/user/auth/referral/model"
	"oggcloudserver/src/user/constants"
	"oggcloudserver/src/user/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//TODO write tests



func CreateReferralModel() (*ref_model.Referral, error) {
	code := make([]byte, 32)
	if _, err := rand.Read(code); err != nil {
		fmt.Fprintf(os.Stderr, "err occurred while random byte generation:\n\t%v\n", err)
		return nil, err
	}
	encodedCode := hex.EncodeToString(code)
	return &ref_model.Referral{
		ID:   uuid.New(),
		Code: encodedCode,
		Used: false,
	}, nil
}

func GetReferralFromCode(code string) (*ref_model.Referral,error) {
	dest := ref_model.Referral{}
	if err := db.DB.Where("code = ?", code).Where("used = ?", false).First(&dest).Error; err != nil {
		return nil,err
	}
	return &dest, nil
}


func CreateReferral(c *gin.Context) {
	email := c.Request.Header.Get(constants.EMAIL_FIELDNAME)
	foundUser, err := model.GetUserFromMail(email)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "err occurred while getting user from mail:\n\t%v\n", err)
		return

	}

	createdReferral, err := CreateReferralModel()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "err occurred while creating referral:\n\t%v\n", err)
		return
	}
	createdReferral.CreatedBy = foundUser.ID

	if err = db.DB.Model(foundUser).Association("Referrals").Append(createdReferral); err != nil {
		c.Status(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "err occurred while getting referral association:\n\t%v\n", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		constants.REFERRAL_CODE_FIELDNAME: createdReferral.Code,
	})

}

func VerifyReferral(c *gin.Context) {
	supposedCode := c.Request.Header.Get(constants.REFERRAL_CODE_FIELDNAME)
	code, verified, err := VerifyReferralImplementation(supposedCode)
	if !verified {
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: \n\t%v\n", err)
		}
		if code != 0 {
			c.Status(code)}
		return
	}  
	
}

func VerifyReferralImplementation(supposedCode string) (int, bool, error)  {
	if supposedCode == "" {
		return http.StatusBadRequest, false, fmt.Errorf("field with name %s not found in the request header", constants.REFERRAL_CODE_FIELDNAME)
	}
	var foundRef ref_model.Referral
	if err := db.DB.Where("code = ?", supposedCode).Where("used = ?", false).First(&foundRef).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusForbidden, false, fmt.Errorf("referral code not found:\n\t%v", err)
		}
		return http.StatusInternalServerError, false,fmt.Errorf("err occurred while getting referral association:\n\t%v", err)
	}
	return 0, true, nil

}
