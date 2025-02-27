package auth_test

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"oggcloudserver/src/db"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/model"
	"oggcloudserver/src/user/testing_material"
	"testing"
	"time"
)

func TestValidity(t *testing.T) {
	assert := assert.New(t)

	testing_material.LoadDotEnv(t)
	testing_material.LoadDB(t)
	defer testing_material.FlushDB()

	var ready_user model.User
	user_id := uuid.New()
	{
		var userjson map[string]interface{}
		marshalledjson, _ := testing_material.GenerateUserJson(t)
		if err := json.Unmarshal(marshalledjson, &userjson); err != nil {
			t.Fatalf("error unmarshalling json:\n\t%v\n", err)
		}

		mdp, err := hex.DecodeString(userjson["password"].(string))
		if !assert.Nil(err) {
			t.Fatalf("error decoding hex to bytes:\n\t%v\n", err)
		}

		bcryptbytes, err := bcrypt.GenerateFromPassword(mdp, bcrypt.DefaultCost)
		if err != nil {
			t.Fatalf("error occured while generating bcrypt pass hash:\n\t%v\n", err)
		}
		mdpbcrypt := string(bcryptbytes)

		ecdhshared := make([]byte, 256)
		_, err = rand.Read(ecdhshared)

		if err != nil {
			t.Fatalf("err occured while reading from random bytes:\n\t%v\n", err)
		}

		ecdhsharedhex := hex.EncodeToString(ecdhshared)

		ready_user = model.User{
			ID:            user_id,
			PasswordHash:  &mdpbcrypt,
			Email:         testing_material.EXAMPLE_MAIL,
			EcdhSharedKey: &ecdhsharedhex,
		}
	}

	if res := db.DB.Create(&ready_user); res.Error != nil {
		t.Fatalf("error inserting instance to database:\n\t%v\n", res.Error)
	}

	{
		myauth, err := auth.CreateInstance(user_id)
		if !assert.Nil(err) {
			t.Fatalf("error creating auth instance:\n\t%v\n", err)
		}
		validcode := myauth.Code

		if err = auth.SaveToDB(myauth); !assert.Nil(err) {
			t.Fatalf("error occured while saving instance to db:\n\t%v\n", err)
		}
		res, err := auth.CheckValidity(validcode)
		if !assert.Nil(err) {
			t.Fatalf("error occured while checking validity:\n\t%v\n", err)
		}
		if !assert.True(res) {
			t.Fatalf("assertion failed: valid code is not valid\n")
		}
	}

	{
		myauth, err := auth.CreateInstance(user_id)
		if !assert.Nil(err) {
			t.Fatalf("error occured while creating auth instance:\n\t%v\n", err)
		}
		myauth.ExpiresAt = time.Now().Add(-5 * time.Minute)
		invalidcode := myauth.Code

		if err = auth.SaveToDB(myauth); !assert.Nil(err) {
			t.Fatalf("error occured while saving instance to db:\n\t%v\n", err)
		}
		res, err := auth.CheckValidity(invalidcode)
		if !assert.Nil(err) {
			t.Fatalf("error occured while checking validity:\n\t%v\n", err)
		}
		if !assert.False(res) {
			t.Fatalf("assertion failed: invalid code is valid\n")
		}

	}

}
