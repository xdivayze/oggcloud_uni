package src

import (
	"oggcloudserver/src/db"
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/file_ops/session"
	"oggcloudserver/src/file_ops/session/Services/retrieve"
	upload "oggcloudserver/src/file_ops/session/Services/upload"
	"oggcloudserver/src/user/auth"
	authmiddleware "oggcloudserver/src/user/auth/auth_middleware"
	"oggcloudserver/src/user/auth/referral"
	ref_model "oggcloudserver/src/user/auth/referral/model"
	"oggcloudserver/src/user/model"
	loginuser "oggcloudserver/src/user/routes/login_user"
	registeruser "oggcloudserver/src/user/routes/register_user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userRegisterRoutes := r.Group("/api/user")
	{
		userRegisterRoutes.POST("/register", registeruser.RegisterUser)
		userRegisterRoutes.POST("/login", loginuser.LoginUser)
		
	}

	r.GET("/api/verify/referral-code", referral.VerifyReferral)
	
	protectedRoutes := r.Group("/", authmiddleware.VerifyCodeMiddleware())
	fileRoutes := protectedRoutes.Group("/api/file")
	{
		fileRoutes.POST("/upload", session.HandleFileUpload)
		fileRoutes.GET("/retrieve", retrieve.HandleRetrieve)
		fileRoutes.GET("/retrieve/get-owner-id", retrieve.GetOwnerFileIDFromPreviewID)
	}
	
	userProtected := protectedRoutes.Group("/api/user/protected")
	{
		userProtected.GET("/create-referral", referral.CreateReferral)
		
	}
	return r
}

func GetDB() (*gorm.DB, error) {
	err := db.Create_DB()
	db.DB.AutoMigrate(&model.User{}, &auth.AuthorizationCode{}, &file.File{}, &upload.Session{}, &ref_model.Referral{})
	return db.DB, err
}
