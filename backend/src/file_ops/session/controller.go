package session

import (
	"log"
	"net/http"
	"oggcloudserver/src/file_ops/session/Services/upload"

	"github.com/gin-gonic/gin"
)
// preview retrievals will work by pulling the last n files
func HandleFileUpload(c *gin.Context) {
	sess, err := upload.HandleFileUploadRequest(c)
	if err != nil {
		log.Printf("error occured while handling file upload request:\n\t%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while processing session request headers"})
		return
	}
	err = upload.HandleFileUpload(c, sess)
	if err != nil {
		log.Printf("error occured while handling file upload:\n\t%v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured while processing files"})
		return
	}

	

	c.JSON(http.StatusCreated, gin.H{"sessionID" : sess.ID.String()})
}