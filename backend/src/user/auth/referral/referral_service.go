package referral

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"oggcloudserver/src/db"
	ref_model "oggcloudserver/src/user/auth/referral/model"
	"oggcloudserver/src/user/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

//TODO write tests

const REFERRAL_CODE_FIELDNAME = "referralCode"

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

func CreateReferral(c *gin.Context) {
	email := c.Request.Header.Get(model.EMAIL_FIELDNAME)
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
		REFERRAL_CODE_FIELDNAME: createdReferral.Code,
	})

}

func VerifyReferral(c *gin.Context) {
	supposedCode := c.Request.Header.Get(REFERRAL_CODE_FIELDNAME)
	if supposedCode == "" {
		c.Status(http.StatusBadRequest)
		fmt.Fprintf(os.Stderr, "field with name %s not found in the request header\n", REFERRAL_CODE_FIELDNAME)
		return
	}

	var foundRef ref_model.Referral
	if err := db.DB.Where("code = ?", supposedCode).First(&foundRef).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusForbidden)
			fmt.Fprintf(os.Stderr, "referral code not found:\n\t%v\n", err)
			return
		}
		c.Status(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "err occurred while getting referral association:\n\t%v\n", err)
		return
	}
	c.Status(http.StatusOK)

}
