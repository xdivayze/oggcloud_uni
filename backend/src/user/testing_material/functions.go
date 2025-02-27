package testing_material

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"oggcloudserver/src"
	"oggcloudserver/src/db"
	"oggcloudserver/src/file_ops/file"
	"oggcloudserver/src/file_ops/session/Services/upload"
	"oggcloudserver/src/oggcrypto"
	"oggcloudserver/src/user/auth"
	"oggcloudserver/src/user/model"
	"testing"

	"github.com/joho/godotenv"
)

const EXAMPLE_MAIL = "example@example.org"
const DOTENV_PATH = "/home/cavej/repositories/oggcloud_dev/backend/.env"

func FlushDB() {
	db.DB.Where("1 = 1").Delete(&model.User{})
	db.DB.Where("1 = 1").Delete(&auth.AuthorizationCode{})
	db.DB.Where("1 = 1").Delete(&upload.Session{})
	db.DB.Where("1 = 1").Delete(&file.File{})
}

func GenerateUserJson(t *testing.T) ([]byte, string) {
	randomBytes := make([]byte, 60)
	_, err := rand.Read(randomBytes)
	if err != nil {
		t.Fatalf("error reading from random buffer:\n\t%v\n", err)
	}
	randomString := hex.EncodeToString(randomBytes)

	_, tp, err := oggcrypto.GenerateECDHPair()
	if err != nil {
		t.Fatalf("error generating ecdh pair:\n\t:%v\n", err)
	}
	pemBlock, err := oggcrypto.EncodePublicKeyToPEM(tp)
	if err != nil {
		t.Fatalf("error encoding public key:\n\t:%v\n", err)
	}

	data, err := json.Marshal(map[string]interface{}{
		model.EMAIL_FIELDNAME:       EXAMPLE_MAIL,
		model.PASSWORD_FIELDNAME:    randomString,
		model.ECDH_PUB_FIELDNAME: pemBlock,
	})

	if err != nil {
		t.Fatalf("error serializing to json:\n\t%v\n", err)
	}
	return data, randomString
}

func LoadDB(t *testing.T) {
	_, err := src.GetDB()
	if err != nil {
		t.Fatalf("error creating database:\n\t%v\n", err)
	}
}

func LoadDotEnv(t *testing.T) {
	err := godotenv.Load(DOTENV_PATH)
	if err != nil {
		t.Fatalf("Error loading .env file %v\n", err)
	}
}
