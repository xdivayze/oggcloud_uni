package constants

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
)

const REFERRAL_CODE_FIELDNAME = "referralCode"
const PASSWORD_FIELDNAME = "password"
const EMAIL_FIELDNAME = "email"
const ECDH_PUB_FIELDNAME = "ecdhPublic"

var BACKEND_DIRECTORY string
var DOTENV_PATH string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	absfilepath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalf("error initializing package, couldn't get absolute file path:%v", err)
	}

	var cwdir = filepath.Dir(absfilepath)
	BACKEND_DIRECTORY = filepath.Join(cwdir, "../../../")
	DOTENV_PATH = fmt.Sprintf("%s/.env", BACKEND_DIRECTORY)
}
