package retrieve

import (
	"fmt"
	"net/http"
	"oggcloudserver/src/db"
	fileops "oggcloudserver/src/file_ops"
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/user/constants"
	"oggcloudserver/src/user/model"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getFileWithOffset(c *gin.Context) (*file.File, error) {

	var foundUser model.User
	var offset int
	var wantPreview bool

	if err := initializeVariables(c, &foundUser, &offset, &wantPreview); err != nil {

		return nil, fmt.Errorf("error occurred while initializing variables:\n\t%w", err)
	}

	foundFile := file.File{}

	if res := db.DB.Where("user_id = ?", foundUser.ID).Where("is_preview = ?", wantPreview).Order("created_at DESC").Offset(offset).Limit(1).Find(&foundFile); res.Error != nil {
		return nil, fmt.Errorf("error finding described user:\n\t%w", res.Error)
	}

	return &foundFile, nil

}



func initializeVariables(c *gin.Context, foundUser *model.User, offset *int, previewMode *bool) error {
	mail := c.Request.Header.Get(constants.EMAIL_FIELDNAME)
	foundUserProto, err := model.GetUserFromMail(mail)
	*foundUser = *foundUserProto
	if err == gorm.ErrRecordNotFound {
		c.JSON(http.StatusNotFound, gin.H{"error": "no user with given email found"})
		return err
	} else if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occurred while associating user with mail"})
		return err
	}

	offsetProto := c.Request.Header.Get(fileops.OFFSET_JSON_FIELDNAME)
	if offsetProto == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "offset field not found in headers"})
	}
	*offset, err = strconv.Atoi(offsetProto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed offset value(should be integer)"})
		return err
	}

	previewModeProto := c.Request.Header.Get(fileops.PREVIEW_WISH_JSON_FIELDNAME)
	if previewModeProto == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wantPreview field not found in headers"})
	}
	*previewMode, err = strconv.ParseBool(previewModeProto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed wantPreview value(should be integer)"})
		return err
	}
	return nil
}
