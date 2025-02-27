package registeruser_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"oggcloudserver/src"
	"oggcloudserver/src/db"
	"oggcloudserver/src/user/model"
	"testing"
	"oggcloudserver/src/user/testing_material"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)
//TODO check for
//faulty e mail



func TestRegisterUser(t *testing.T) {
	testing_material.LoadDotEnv(t)
	testing_material.LoadDB(t)
	defer db.DB.Where("1 = 1").Delete(&model.User{})

	gin.SetMode(gin.TestMode)
	r := src.SetupRouter()
	w := httptest.NewRecorder()

	data, _ := testing_material.GenerateUserJson(t)

	endpoint := "/api/user/register"
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("error creating new request:\n\t%v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d\n\tjsonBody:%s", w.Code, w.Body.String())
	}
	t.Logf("responseBody:\n\t%s\n", w.Body.String())
	_, res := model.GetUserFromMail(testing_material.EXAMPLE_MAIL)
	if res != nil {
		t.Fatalf("error occured while getting user from database:\n\t%v\n", res.Error())
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &jsonData); err != nil {
		t.Logf("error unmarshaling json:\n\t%v\n", err)
	}

	id, exists := jsonData["id"]
	if !exists {
		t.Logf("ID field doesn't exist on return json")
	}
	uuuid, err := uuid.Parse(id.(string))
	if err != nil {
		t.Logf("couldn't parse uuid:\n\t%v\n", err)
	}
	_, res = model.GetUserFromID(uuuid)
	if res != nil {
		t.Fatalf("error occured while getting user from database by ID:\n\t%v\n", res.Error())
	}
	_, res = model.GetUserFromMail(testing_material.EXAMPLE_MAIL)
	if res != nil {
		t.Fatalf("error occured while getting user from database:\n\t%v\n", res.Error())
	}
}
