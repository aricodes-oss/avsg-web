package web

import (
	. "avsg/embeds"
	_ "embed"

	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	router = NewRouter()
	os.Exit(m.Run())
}

func TestEndpoints(t *testing.T) {
	var tests = []struct {
		endpoint string
		sample   []byte
	}{
		{"/api/encrypt", Samples.Decrypted},
		{"/api/decrypt", Samples.Encrypted},
	}

	for _, test := range tests {
		t.Run(test.endpoint, func(t *testing.T) {
			assert := assert.New(t)
			w := httptest.NewRecorder()

			req, err := newFileUploadRequest(test.endpoint, test.sample)
			assert.Nil(err)

			router.ServeHTTP(w, req)
			if w.Result().StatusCode != http.StatusOK {
				t.Fatal(w.Body.String())
			}
			assert.Equal(w.Result().StatusCode, http.StatusOK)
		})
	}
}

func newFileUploadRequest(uri string, data []byte) (*http.Request, error) {
	file := bytes.NewReader(data)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "Save0.sav")
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	return req, nil
}
