package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const DIRECTORY_BASE = "/home/cavej/repositories/oggcloud_dev/backend/Storage/Files" //TODO fix value absolute path 

var DirectorySession string

func HandleFileUpload(c *gin.Context, session *Session) error {
	log.SetPrefix("ERR: ")
	file, _, err := c.Request.FormFile("file") //TODO change this stuff with a multipart reader https://gist.github.com/ZenGround0/49e4a1aa126736f966a1dfdcb84abdae 
	if err != nil {
		return fmt.Errorf("error occurred while retrieving file from form:\n\t%w", err)
	}
	defer file.Close()

	DirectorySession = fmt.Sprintf("%s/%s/%s", DIRECTORY_BASE, session.UserID, session.ID)
	if err = os.MkdirAll(DirectorySession, 0777); err != nil {
		return fmt.Errorf("error occurred while creating directory at path:%s:\n\t%w", DirectorySession, err)
	}
	if err = extractTarGz(file, session); err != nil {
		return fmt.Errorf("error occurred while extracting from tarball:\n\t%w", err)
	}
	return nil

}
