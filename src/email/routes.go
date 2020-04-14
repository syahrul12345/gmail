package email

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func verificationEmailHandler(c *gin.Context) {
	/***
		Payload only contains verification email with an authenticator hash in the authorization header. This will be handled by the auth middleware
		{
			"email": "muhdsyahrulnizam123@gmail.com"
		}
	***/
	verificationReq := &VerificationRequest{}
	err := c.ShouldBindWith(verificationReq, binding.JSON)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	email := verificationReq.Email
	token, err := verificationReq.CreateToken()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	err = send(email, token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
