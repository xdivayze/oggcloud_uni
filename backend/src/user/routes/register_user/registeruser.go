package register_user

import (
	"fmt"
	"log"
	"net/http"
	"oggcloudserver/src/db"
	"oggcloudserver/src/functions"
	"oggcloudserver/src/user/constants"
	"oggcloudserver/src/user/model"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// TODO email check
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
	fieldMap[constants.EMAIL_FIELDNAME] = &mail
	fieldMap[constants.PASSWORD_FIELDNAME] = &passwordHex
	fieldMap[constants.ECDH_PUB_FIELDNAME] = &ecdhClientPub
	fieldMap[constants.REFERRAL_CODE_FIELDNAME] = &referralCode

	s := functions.DoFieldAssign(c, jsonData, fieldMap)
	if s != 0 {
		log.Printf("error doing field assignments, returned:%d", s)
		return
	}
	id := uuid.New()
	if res, err := processReferral(referralCode, id, c); !res {
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

	sharedKey, serverPub, err := model.GenerateAndEncryptSharedKey(ecdhClientPub) //salt is prepended to shared key
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
