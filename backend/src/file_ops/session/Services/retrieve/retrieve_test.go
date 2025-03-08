package retrieve_test

import (
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"oggcloudserver/src"
	fileops "oggcloudserver/src/file_ops"
	"oggcloudserver/src/oggcrypto"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/constants"
	"oggcloudserver/src/user/testing_material"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestDownloadIntegrityID(t *testing.T) {
	require := require.New(t)
	testing_material.ModeFlush = false
	testing_material.TestDataHandling(t)
	//panic("a")
	testing_material.ModeFlush = true

	defer func() {
		if testing_material.ModeFlush {
			testing_material.FlushDB()
			os.RemoveAll(testing_material.UDir)
		}
	}()

	endpoint := "/api/file/retrieve"

	gin.SetMode(gin.TestMode)
	r := src.SetupRouter()

	var id string
	{ //get file id
		req, err := http.NewRequest("GET", endpoint, nil)
		require.Nil(err)

		req.Header.Set(constants.EMAIL_FIELDNAME, testing_material.EXAMPLE_MAIL)
		req.Header.Set(auth.AUTH_CODE_FIELDNAME, testing_material.Auth)
		req.Header.Set(fileops.PULL_METHOD_JSON_FIELDNAME, "offset")
		req.Header.Set(fileops.PREVIEW_WISH_JSON_FIELDNAME, "true")
		req.Header.Set(fileops.OFFSET_JSON_FIELDNAME, "0")

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		require.Equal(w.Code, http.StatusOK)

		multipartValMap := make(map[string]string)
		var calculatedSum string

		reader := multipart.NewReader(w.Body, w.Header().Get("Content-Type")[len("multipart/form-data; boundary="):])
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			require.Nil(err)

			formName := part.FormName()
			fileName := part.FileName()
			if fileName != "" {
				var err error
				calculatedSum, err = oggcrypto.CalculateSHA256sum(part)
				require.Nil(err)

			} else if formName != "" {
				val, err := io.ReadAll(part)
				require.Nil(err)
				multipartValMap[formName] = string(val)
			}

			part.Close()

		}

		require.Equal(calculatedSum, multipartValMap["checksum"])
		id = multipartValMap[fileops.FILE_ID_JSON_FIELDNAME]
	}

	require.NotEmpty(id)

	{ //get file from id
		req, err := http.NewRequest("GET", endpoint, nil)
		require.Nil(err)

		req.Header.Set(constants.EMAIL_FIELDNAME, testing_material.EXAMPLE_MAIL)
		req.Header.Set(auth.AUTH_CODE_FIELDNAME, testing_material.Auth)
		req.Header.Set(fileops.PULL_METHOD_JSON_FIELDNAME, "id")
		req.Header.Set(fileops.PREVIEW_WISH_JSON_FIELDNAME, "true")
		req.Header.Set(fileops.FILE_ID_JSON_FIELDNAME, id)

		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		require.Equal(w.Code, http.StatusOK)

		multipartValMap := make(map[string]string)
		var calculatedSum string

		reader := multipart.NewReader(w.Body, w.Header().Get("Content-Type")[len("multipart/form-data; boundary="):])
		for {
			part, err := reader.NextPart()
			if err == io.EOF {
				break
			}
			require.Nil(err)

			formName := part.FormName()
			fileName := part.FileName()
			if fileName != "" {
				var err error
				calculatedSum, err = oggcrypto.CalculateSHA256sum(part)
				require.Nil(err)

			} else if formName != "" {
				val, err := io.ReadAll(part)
				require.Nil(err)
				multipartValMap[formName] = string(val)
			}

			part.Close()

		}

		require.Equal(calculatedSum, multipartValMap["checksum"])
	}

}

func TestDownloadIntegrity(t *testing.T) {
	require := require.New(t)

	testing_material.ModeFlush = false
	testing_material.TestDataHandling(t)
	//panic("a")
	testing_material.ModeFlush = true

	defer func() {
		if testing_material.ModeFlush {
			testing_material.FlushDB()
			os.RemoveAll(testing_material.UDir)
		}
	}()

	endpoint := "/api/file/retrieve"

	gin.SetMode(gin.TestMode)
	r := src.SetupRouter()
	req, err := http.NewRequest("GET", endpoint, nil)
	require.Nil(err)

	req.Header.Set(constants.EMAIL_FIELDNAME, testing_material.EXAMPLE_MAIL)
	req.Header.Set(auth.AUTH_CODE_FIELDNAME, testing_material.Auth)
	req.Header.Set(fileops.PULL_METHOD_JSON_FIELDNAME, "offset")
	req.Header.Set(fileops.PREVIEW_WISH_JSON_FIELDNAME, "true")
	req.Header.Set(fileops.OFFSET_JSON_FIELDNAME, "0")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(w.Code, http.StatusOK)

	multipartValMap := make(map[string]string)
	var calculatedSum string

	reader := multipart.NewReader(w.Body, w.Header().Get("Content-Type")[len("multipart/form-data; boundary="):])
	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		require.Nil(err)

		formName := part.FormName()
		fileName := part.FileName()
		if fileName != "" {
			var err error
			calculatedSum, err = oggcrypto.CalculateSHA256sum(part)
			require.Nil(err)

		} else if formName != "" {
			val, err := io.ReadAll(part)
			require.Nil(err)
			multipartValMap[formName] = string(val)
		}

		part.Close()

	}

	require.Equal(calculatedSum, multipartValMap["checksum"])

}
