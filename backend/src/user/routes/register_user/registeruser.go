package registeruser

import (
	"fmt"
	"log"
	"net/http"
	"oggcloudserver/src/db"
	"oggcloudserver/src/functions"
	"oggcloudserver/src/user/auth/referral"
	"oggcloudserver/src/user/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO email check
// TODO implement invitation code
func RegisterUser(c *gin.Context) {
	log.SetPrefix("ERROR: ")
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("couldn't parse json\n\t:%v", err.Error())})
		return
	}

	var mail string
	var passwordHex string
	var ecdhClientPub string
	var referralCode string

	fieldMap := make(map[string]interface{})
	fieldMap[model.EMAIL_FIELDNAME] = &mail
	fieldMap[model.PASSWORD_FIELDNAME] = &passwordHex
	fieldMap[model.ECDH_PUB_FIELDNAME] = &ecdhClientPub
	fieldMap[referral.REFERRAL_CODE_FIELDNAME] = &referralCode

	s := functions.DoFieldAssign(c, jsonData, fieldMap)
	if s != 0 {
		log.Printf("error doing field assignments, returned:%d", s)
		return
	}
	id := uuid.New()
	if res, err := processReferral(referralCode, id,c); !res  { //TODO add referral to tests
		if err != nil {
			fmt.Fprintf(os.Stderr, "error occurred while processing referral:\n\t%v\n", err)
			return
		}
		return
	}

	password, err := processPassword(c, passwordHex)
	if err != nil {
		log.Printf("error occurred while processing password: \n\t%v\n", err)
		return
	}

	sharedKey, serverPub, err := model.GenerateAndEncryptSharedKey(ecdhClientPub) //salt is prepended to sharedkey
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred registering user"})
		log.Printf("error occurred while generating and encrypting the shared key:\n\t%v\n", err)
		return
	}
	
	user := model.User{
		ID:            id,
		Email:         mail,
		PasswordHash:  &password,
		EcdhSharedKey: &sharedKey,
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred registering user"})
		log.Printf("error occurred while registering user to database:\n\t%v\n", result.Error)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":                  id.String(),
		"ServerECDHPublicKey": serverPub,
	})
	log.SetPrefix("INFO: ")
	log.Println("user created:\n", user.ToString())

}
