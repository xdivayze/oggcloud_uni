package retrieve

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"oggcloudserver/src/db"
	fileops "oggcloudserver/src/file_ops"
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/user/constants"
	"oggcloudserver/src/user/model"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func GetOwnerFileIDFromPreviewID(c *gin.Context) { 
	foundFile := &file.File{}
	foundUser := &model.User{}

	email := c.Request.Header.Get(constants.EMAIL_FIELDNAME)
	if err := db.DB.Where("email = ?", email).Find(foundUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		fmt.Fprintf(os.Stderr, "error occurred while finding user with email %s:\n\t%v\n", email, err)
		c.Status(http.StatusInternalServerError)
		return
	}
	prID := c.GetHeader(fileops.FILE_ID_JSON_FIELDNAME)
	if prID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "fileID field doesn't exist in the header"})
		return
	}
	if err := db.DB.Where("preview_id = ?", prID).Where("user_id = ?", foundUser.ID).Find(&foundFile).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		c.Status(http.StatusInternalServerError)
		fmt.Fprintf(os.Stderr, "error occurred while finding file with preview_id %s\n\t%v\n", prID, err)
		return
	}
	c.JSON(http.StatusFound, gin.H{"id": foundFile.ID})
}

func HandleRetrieve(c *gin.Context) { //work with offset or ID get requests, not multiple photos at one request
	log.SetPrefix("ERR: ")
	returnedFile := &file.File{}
	if c.Request.Header.Get(fileops.PULL_METHOD_JSON_FIELDNAME) == "offset" {
		var err error
		returnedFile, err = getFileWithOffset(c)

		if err != nil {
			log.Printf("error occurred while getting image:\n\t%v\n", err)
			return
		}
	} else if c.Request.Header.Get(fileops.PULL_METHOD_JSON_FIELDNAME) == "id" {
		fileID := c.Request.Header.Get(fileops.FILE_ID_JSON_FIELDNAME)
		if fileID == "" {
			log.Printf("error occurred while getting request header with field %s\n", fileops.FILE_ID_JSON_FIELDNAME)
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("field with name %s doesn't exist", fileops.FILE_ID_JSON_FIELDNAME)})
			return
		}
		if res := db.DB.Find(returnedFile, "id = ?", fileID); res.Error != nil {
			log.Printf("error occurred while finding file with id %s:\n\t%v\n", fileID, res.Error)
			if res.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("file with id %s doesn't exist", fileID)})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while trying to find file"})
			return
		}
		u := &model.User{}
		db.DB.Where("email = ?", c.GetHeader("email")).Find(u)
		if returnedFile.UserID != u.ID {
			log.Printf("user doesn't own requested file")
			c.Status(http.StatusForbidden)
			return
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "pull method not specified"})
		return
	}

	if err := doLoadFileAndStream(c, returnedFile); err != nil {
		log.Printf("error occurred while loading and streaming file:\n\t%v\n ", err)
		return
	}
}
