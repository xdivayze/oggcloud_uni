package upload

import (
	"fmt"
	"oggcloudserver/src/db"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleFileUploadRequest(c *gin.Context) (*Session, error) {
	id := uuid.New()

	file_num, err := strconv.Atoi(c.Request.FormValue("file_count"))
	if err != nil {
		return nil, fmt.Errorf("error occured while parsing to int")

	}
	uid, err := uuid.Parse(c.Request.FormValue("id"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse form uid to uuid:\n\t%w", err)
	}
	session_key := c.Request.FormValue("session_key")
	if session_key == "" {
		return nil, fmt.Errorf("session_key field doesn't exist")
	}
	checksum := c.Request.FormValue("checksum")
	if checksum == "" {
		return nil, fmt.Errorf("checksum field doesn't exist")
	}

	current_session := Session{
		ID:              id,
		SessionKey:      session_key,
		FileNumber:      file_num,
		UserID:          uid,
		SessionChecksum: checksum,
	}

	if res := db.DB.Save(&current_session); res.Error != nil {
		return nil, fmt.Errorf("error occured while saving instance to DB:\n\t%w", res.Error)
	}
	return &current_session, nil

}
