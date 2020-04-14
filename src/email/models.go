package email

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// VerificationRequest holds information to send a verification email to.
type VerificationRequest struct {
	Email string `json:"email"`
}

// Token is a struct that holds the JWT object
type Token struct {
	Email string
	jwt.StandardClaims
}

// CreateToken will create a JWT token for the VerificationRequest
func (v *VerificationRequest) CreateToken() (string, error) {
	tk := &Token{v.Email, jwt.StandardClaims{
		IssuedAt: time.Now().Unix(),
		Issuer:   "VerificationService",
		Subject:  "Authtoken",
	}}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	// Our standard claims is stored in the db.Token struct
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
