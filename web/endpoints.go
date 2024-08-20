package web

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"avsg/crypto"
)

func Encrypt(c *gin.Context) {
	// Retrieve the uploaded file
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		// c.Redirect(http.StatusTemporaryRedirect, "/error") // TODO: Make error page
		return
	}

	// Get a handle to it, read it into memory
	handle, _ := file.Open()
	defer handle.Close()
	contents, _ := io.ReadAll(handle)

	encrypted, err := crypto.Encrypt(contents)
	if err != nil {
		c.Redirect(http.StatusFound, "/error")
	}

	// Respond directly with the encrypted save file
	c.Data(200, "application/octet-stream", encrypted)
}

func Decrypt(c *gin.Context) {
	// Retrieve the uploaded file
	file, err := c.FormFile("file")
	if err != nil || file == nil {
		c.Redirect(http.StatusFound, "/error") // TODO: Make error page
		return
	}

	// Get a handle to it, read it into memory
	handle, _ := file.Open()
	defer handle.Close()
	contents, _ := io.ReadAll(handle)

	decrypted, err := crypto.Decrypt(contents)
	if err != nil {
		c.Redirect(http.StatusTemporaryRedirect, "/error")
	}

	// Respond directly with the decrypted save file
	c.Data(200, "application/octet-stream", decrypted)
}
