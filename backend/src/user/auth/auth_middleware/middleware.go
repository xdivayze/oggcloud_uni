package authmiddleware

import (
	"net/http"
	"oggcloudserver/src/db"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/model"

	"github.com/gin-gonic/gin"
)

func VerifyCodeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Request.Header.Get(model.EMAIL_FIELDNAME)
		if email == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "email field doesn't exist in header",
			})
			c.Abort()
			return
		}
		authCode := c.Request.Header.Get(auth.AUTH_CODE_FIELDNAME)
		if authCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "authCode field doesn't exist in header",
			})
			c.Abort()
			return
		}
		var u model.User
		if res := db.DB.Where("Email = ?", email).Find(&u); res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "no user with given email exists",
			})
			c.Abort()
			return
		}
		var authCodes []auth.AuthorizationCode
		db.DB.Model(&u).Association("AuthorizationCodes").Find(&authCodes)
		for _, code := range authCodes {
			if (authCode == code.Code) && code.IsValid(true) {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error":"authorization code doesn't exist or is invalid"})

	}
}
