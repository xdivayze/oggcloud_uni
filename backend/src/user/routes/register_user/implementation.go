package registeruser

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const PASSWORD_LENGTH = 9

var ErrPasswordTooLong = fmt.Errorf("password length exceeds %d characters", PASSWORD_LENGTH)

func processPassword(c *gin.Context, passwordhex string) (string, error) {
	hexpass, err := hex.DecodeString(passwordhex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured registering user"})
		return "", fmt.Errorf("error occured while generating bytes from hex:\n\t%w", err)
	}

	if len(hexpass) > PASSWORD_LENGTH*8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": ErrPasswordTooLong.Error()})
		return "", ErrPasswordTooLong
	}

	bcryptPass, err := bcrypt.GenerateFromPassword(hexpass, bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error occured registering user"})
		return "", fmt.Errorf("error occured while generating bcrypt:\n\t%w", err)
	}
	return string(bcryptPass), nil

}
