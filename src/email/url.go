package email

import "github.com/gin-gonic/gin"

func Register(r *gin.RouterGroup) {
	r.POST("/sendVerificationEmail", verificationEmailHandler)
}
