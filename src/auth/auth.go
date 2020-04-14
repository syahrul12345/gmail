package auth

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		secretKey := os.Getenv("INTERNAL_SECRET_KEY")
		apiKey := c.GetHeader("Authorization")

		apiKeyArray := strings.Split(apiKey, " ")
		token := apiKeyArray[1]

		fullPath := c.FullPath()
		// rawData, _ := c.GetRawData()
		// We cannot use GetRawData as it reads directly from the stream.
		// Copy it into a buffer first!
		bodyCopy := new(bytes.Buffer)
		_, err := io.Copy(bodyCopy, c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		rawData := bodyCopy.Bytes()
		rawDataString := string(rawData)
		// Putback a new reader into the request body
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(rawData))

		dataList := []string{fullPath, rawDataString, secretKey}
		concatonatedData := strings.Join(dataList, "")

		sha := sha256.Sum256([]byte(concatonatedData))
		shaString := hex.EncodeToString(sha[:])
		log.Println(shaString)
		if shaString != token {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid api key."})
			c.Abort()
			return
		}

		c.Next()
	}
}
