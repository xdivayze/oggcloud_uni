package testing_material

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"oggcloudserver/src"
	"oggcloudserver/src/db"
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/file_ops/session/Services/upload"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/model"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

const TEST_TAR = "/home/cavej/repositories/oggcloud_dev/backend/Storage/testing/uploadtest/test.tar.gz"

var ModeFlush = true
var Udir string
var Auth string

func TestDBIntegrity(t *testing.T) {
	require := require.New(t)
	ModeFlush = false
	TestDataHandling(t)
	ModeFlush = true
	defer func() {
		if ModeFlush {
			FlushDB()
			os.RemoveAll(Udir)
		}
	}()
	lx := strings.Split(Udir, "/")
	id := lx[len(lx)-1]
	var u model.User
	var l []upload.Session
	uid, err := uuid.Parse(id)
	require.Nil(err)
	res := db.DB.Find(&u, uid)
	require.Nil(res.Error)

	err = db.DB.Model(&u).Association("Sessions").Find(&l)
	require.Nil(err)

	storageDir, err := os.ReadDir(fmt.Sprintf("%s/%s/%s", Udir, l[0].ID, "Storage"))
	require.Nil(err)
	for _, f := range storageDir {
		var foundFile file.File
		res := db.DB.Where("file_name = ?", f.Name()).First(&foundFile)
		require.Nil(res.Error)
		if !strings.HasSuffix(foundFile.FileName, "json") {
			require.True(foundFile.HasPreview || foundFile.IsPreview)
			if foundFile.HasPreview {
				var previewFile file.File
				db.DB.Model(&foundFile).Association("Preview").Find(&previewFile)
				require.NotNil(previewFile)
			}

		}
	}
}

func TestDataHandling(t *testing.T) {
	LoadDotEnv(t)
	LoadDB(t)

	defer func() {
		if ModeFlush {
			FlushDB()
		}
	}()

	gin.SetMode(gin.TestMode)
	r := src.SetupRouter()

	id, authcode := DoCreateUser(t, r)
	Udir = fmt.Sprintf("%s/%s", upload.DIRECTORY_BASE, id.String())

	defer func() {
		if ModeFlush {
			os.RemoveAll(Udir)
		}
	}()

	file, err := os.Open(TEST_TAR)
	if err != nil {
		t.Fatalf("error trying to open test tarball:\n\t%v\n", err)
	}

	defer file.Close()

	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)

	filepart, err := writer.CreateFormFile("file", "mytar.tar.gz")
	if err != nil {
		t.Fatalf("error creating form file:\n\t%v\n", err)
	}
	if _, err = io.Copy(filepart, file); err != nil {
		t.Fatalf("error with io operation:\n\t%v\n", err)
	}

	if err = writer.WriteField("id", id.String()); err != nil {
		t.Fatalf("error occured while writing field")
	}
	if err = writer.WriteField("file_count", "2"); err != nil {
		t.Fatalf("error occured while writing field")
	}
	if err = writer.WriteField("checksum", "d7a51f12f8a85e315936d09acd74daed245551bcb77e450c88c8a05179ddf96b"); err != nil {
		t.Fatalf("error occured while writing field")
	}

	ra := make([]byte, 64)
	if _, err = rand.Read(ra); err != nil {
		t.Fatalf("error generating random values:\n\t%v", err)
	}

	if err = writer.WriteField("session_key", hex.EncodeToString(ra)); err != nil {
		t.Fatalf("error occured while writing field")
	}
	writer.Close()

	req, err := http.NewRequest("POST", "/api/file/upload", &requestBody)
	if err != nil {
		t.Fatalf("error generating new request:\n\t%v\n", err)
	}

	req.Header.Set(model.EMAIL_FIELDNAME, EXAMPLE_MAIL)
	req.Header.Set(auth.AUTH_CODE_FIELDNAME, authcode)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("status returned isnt 201 but %d", w.Code)
	}

	var unmarshaled map[string]interface{}
	if err = json.Unmarshal(w.Body.Bytes(), &unmarshaled); err != nil {
		t.Fatalf("error occured while unmarshalling:\n\t%v\n", err)
	}
	sid, err := uuid.Parse(unmarshaled["sessionID"].(string))
	if err != nil {
		t.Fatalf("error occured while parsing to uuid:\n\t%v\n", err)
	}

	require.DirExists(t, fmt.Sprintf("%s/%s/Storage", Udir, sid))
	require.DirExists(t, fmt.Sprintf("%s/%s/Preview", Udir, sid))

}

func DoCreateUser(t *testing.T, r *gin.Engine) (uuid.UUID, string) {
	userjson, password := GenerateUserJson(t)
	w := httptest.NewRecorder()
	endpoint := "/api/user/register"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(userjson))
	if err != nil {
		t.Fatalf("error creating new request:\n\t%v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d\n\tjsonBody:%s", w.Code, w.Body.String())
	}

	var jsonobj map[string]interface{}
	if err = json.Unmarshal(w.Body.Bytes(), &jsonobj); err != nil {
		t.Fatalf("error occured while unmarshalling:\n\t%v\n", err)
	}
	id, err := uuid.Parse(jsonobj["id"].(string))
	if err != nil {
		t.Fatalf("error occured while parsing to uuid:\n\t%v\n", err)
	}
	return id, DoLogin(t, password, r)
}

func DoLogin(t *testing.T, password string, r *gin.Engine) string {
	jsonMap := map[string]interface{}{
		model.EMAIL_FIELDNAME:    EXAMPLE_MAIL,
		model.PASSWORD_FIELDNAME: password,
	}

	jsonBytes, err := json.Marshal(jsonMap)
	require.Nil(t, err)

	endpoint := "/api/user/login"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBytes))
	require.Nil(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &responseBody)
	require.Nil(t, err)

	auth, exists := responseBody[auth.AUTH_CODE_FIELDNAME]
	authl := auth.(string)
	require.True(t, exists)
	Auth = authl
	return authl

}
