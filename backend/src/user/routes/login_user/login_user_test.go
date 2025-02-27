package loginuser_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"oggcloudserver/src"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/model"
	loginuser "oggcloudserver/src/user/routes/login_user"
	"oggcloudserver/src/user/testing_material"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func generateLoginJson(passwd string,t *testing.T)[]byte {
	data , err := json.Marshal(map[string]interface{}{
		"email":testing_material.EXAMPLE_MAIL,
		"password":passwd,
	})
	if err != nil {
		t.Fatalf("error while generating login json:\n\t%v\n", err)
	}
	return data
}

func TestLogin(t *testing.T) {
	testing_material.LoadDotEnv(t)
	testing_material.LoadDB(t)

	defer testing_material.FlushDB()

	gin.SetMode(gin.TestMode)
	r := src.SetupRouter()
	
	userjson, passwd := testing_material.GenerateUserJson(t)
	{
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
	}
	var loginreturn map[string]interface{}
	{
		w := httptest.NewRecorder()
		loginjson := generateLoginJson(passwd, t)
		endpoint := "/api/user/login"
		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(loginjson))
		if err != nil {
			t.Fatalf("error creating new request:\n\t%v\n", err)
		}
		req.Header.Set("Content-Type", "application/json")
		r.ServeHTTP(w, req)
		if w.Code != http.StatusOK {
			t.Fatalf("expected 200, got %d\n\tjsonBody:%s", w.Code, w.Body.String())
		}

		if err = json.Unmarshal(w.Body.Bytes(), &loginreturn); err != nil {
			t.Fatalf("error unmarshalling json:\n\t%v\n", err)
		}
		t.Logf("success, body returned: \n\t%s", w.Body.String())

	}
	code, s := loginreturn["authCode"].(string)
	if !s {
		t.Fatalf("field authCode doesn't exist on returned json")
	}
	foundAuth, err := auth.RetrieveFromDB(code)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("auth instance not found:\n\t%s\n", code)
	} else if err != nil {
		t.Fatalf("error occured while retrieving instance:\n\t%v\n", err)
	}
	ownerID := foundAuth.UserID
	foundUser, err := model.GetUserFromID(ownerID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		t.Fatalf("user instance not found:\n\t%s\n", code)
	} else if err != nil {
		t.Fatalf("error occured while retrieving instance:\n\t%v\n", err)
	}

	if err = loginuser.CheckPasswordHash(passwd, foundUser); err != nil {
		log.Fatalf("passwords don't match: \n\t%v\n", err)
	}
}
