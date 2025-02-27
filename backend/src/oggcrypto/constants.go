package oggcrypto

import (
	"log"
	"path/filepath"
	"runtime"
)

const AES_KEY_SIZE = 32
const SALT_LENGTH = 16
const NONCE_LENGTH = 12

var cwdir string
var data_folder string
var MASTERKEY_PATH string
var ECDHKEY_PATH string

func init() {
	_, filename, _, _ := runtime.Caller(0)
	absfilepath, err := filepath.Abs(filename)
	if err != nil {
		log.Fatalf("error initializing package, couldn't get absolute file path:%v", err)
	}

	cwdir = filepath.Dir(absfilepath)
	data_folder = filepath.Join(cwdir, "data")
	MASTERKEY_PATH = filepath.Join(data_folder, "masterkey.hex")
	ECDHKEY_PATH = filepath.Join(data_folder, "ecdhkeys")

}
