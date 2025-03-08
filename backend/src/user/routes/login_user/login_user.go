package loginuser

import (
	"fmt"
	"log"
	"net/http"
	"oggcloudserver/src/functions"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/constants"
	"oggcloudserver/src/user/model"

	"github.com/gin-gonic/gin"
)


func LoginUser(c *gin.Context) {
	log.SetPrefix("ERROR: ")
	var jsonData map[string]interface{}

	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("couldn't parse json\n\t:%v", err.Error())})
		return
	}

	var passwordHex string
	var email string

	fieldMap := make(map[string]interface{})
	fieldMap[constants.EMAIL_FIELDNAME] = &email
	fieldMap[constants.PASSWORD_FIELDNAME] = &passwordHex
	s := functions.DoFieldAssign(c, jsonData, fieldMap)
	if s != 0 {
		log.Printf("error doing field assignments, returned:%d", s)
		return
	}

	foundUser, err := model.GetUserFromMail(email)
	if err != nil {
		log.Printf("error occurred while querying database:\n\t%v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user not found with mail %s", email)})
		return
	}

	if err = CheckPasswordHash(passwordHex, foundUser); err != nil {
		log.Printf("error occurred while checking password hash:\n\t%v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password doesn't match"})
		return
	}

	code, err := auth.CreateInstance(foundUser.ID)
	if err != nil {
		log.Printf("error occurred while creating auth instance:\n\t%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while logging in"})
		return
	}
	if err = auth.SaveToDB(code); err != nil {
		log.Printf("error occurred while saving auth instance:\n\t%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while logging in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{auth.AUTH_CODE_FIELDNAME: code.Code, "expiresAt": code.ExpiresAt})

}
