package web

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"avsg/crypto"
)

var (
	Encrypt = makeEndpoint(crypto.Encrypt, ".xml", ".sav")
	Decrypt = makeEndpoint(crypto.Decrypt, ".sav", ".xml")
)

func makeEndpoint(cryptFunc func([]byte) ([]byte, error), oldExt, newExt string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve the uploaded file
		file, err := c.FormFile("file")
		if err != nil || file == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err, "file": file})
			return
		}

		// Get a handle to it, read it into memory
		handle, _ := file.Open()
		defer handle.Close()
		contents, _ := io.ReadAll(handle)

		transformed, err := cryptFunc(contents)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// Respond directly with the altered save file
		c.Header("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, swapExt(file.Filename, oldExt, newExt)))
		c.Data(200, "application/octet-stream", transformed)
	}
}

func swapExt(str, old, new string) string {
	name, _ := strings.CutSuffix(str, old)
	return name + new
}
