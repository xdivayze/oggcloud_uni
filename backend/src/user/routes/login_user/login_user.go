package loginuser

import (
	"fmt"
	"log"
	"net/http"
	"oggcloudserver/src/functions"
	"oggcloudserver/src/user/auth"
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

	var passwordhex string
	var email string

	fieldmap := make(map[string]interface{})
	fieldmap[model.EMAIL_FIELDNAME] = &email
	fieldmap[model.PASSWORD_FIELDNAME] = &passwordhex
	s := functions.DoFieldAssign(c, jsonData, fieldmap)
	if s != 0 {
		log.Printf("error doing field assignments, returned:%d", s)
		return
	}

	foundUser, err := model.GetUserFromMail(email)
	if err != nil {
		log.Printf("error occured while querying database:\n\t%v\n", err)
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("user not found with mail %s", email)})
		return
	}

	if err = CheckPasswordHash(passwordhex, foundUser); err != nil {
		log.Printf("error occured while checking password hash:\n\t%v\n", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "password doesn't match"})
		return
	}

	code, err := auth.CreateInstance(foundUser.ID)
	if err != nil {
		log.Printf("error occured while creating auth instance:\n\t%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while logging in"})
		return
	}
	if err = auth.SaveToDB(code); err != nil {
		log.Printf("error occured while saving auth instance:\n\t%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while logging in"})
		return
	}

	c.JSON(http.StatusOK, gin.H{auth.AUTH_CODE_FIELDNAME: code.Code, "expiresAt": code.ExpiresAt})

}
