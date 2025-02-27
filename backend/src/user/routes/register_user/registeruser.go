package registeruser

import (
	"fmt"
	"log"
	"net/http"
	"oggcloudserver/src/db"
	"oggcloudserver/src/functions"
	"oggcloudserver/src/user/model"

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
	var passwordhex string
	var ecdhclientpub string

	fieldmap := make(map[string]interface{})
	fieldmap[model.EMAIL_FIELDNAME] = &mail
	fieldmap[model.PASSWORD_FIELDNAME] = &passwordhex
	fieldmap[model.ECDH_PUB_FIELDNAME] = &ecdhclientpub

	s := functions.DoFieldAssign(c, jsonData, fieldmap)
	if s != 0 {
		log.Printf("error doing field assignments, returned:%d", s)
		return
	}

	password, err := processPassword(c, passwordhex)
	if err != nil {
		log.Printf("error occured while processing password: \n\t%v\n", err)
		return
	}

	sharedkey, serverpub, err := model.GenerateAndEncryptSharedKey(ecdhclientpub) //salt is prepended to sharedkey
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured registering user"})
		log.Printf("error occured while generating and encrypting the shared key:\n\t%v\n", err)
		return
	}
	id := uuid.New()
	user := model.User{
		ID:            id,
		Email:         mail,
		PasswordHash:  &password,
		EcdhSharedKey: &sharedkey,
	}
	result := db.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured registering user"})
		log.Printf("error occured while registering user to database:\n\t%v\n", result.Error)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"id":                  id.String(),
		"ServerECDHPublicKey": serverpub,
	})
	log.SetPrefix("INFO: ")
	log.Println("user created:\n", user.ToString())

}
