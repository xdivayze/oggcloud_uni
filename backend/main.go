package main

import (
	"fmt"
	"log"
	"oggcloudserver/src"
	"oggcloudserver/src/oggcrypto"
	"oggcloudserver/src/user"
	"oggcloudserver/src/user/testing_material"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotenv() error {
	return godotenv.Load()
}
//TODO fix absolute paths throughout the codebase
//TODO request a demo referral, the user will be destroyed within 2 hours of registration
// user access levels????
// forward requests through a dpi tunnel to surpass censorship
// implement rate limiting
// implement human test
// -> select apo,onur,uygar,isra etc's head, leg etc
// feature to share files between many users
// create a queue middleware for requests concerning other users with structs in db
// links from client to other clients through server using e2ee to share pictures( A sends B his key through
// server so that they can both read the file) DO THIS LATER
// When session groups are being created an option to send other users a request to view session??
// use ecdh for multiple connections to derive a shared key
// maybe some kind of a config file to load sessions for client
//TODO SHARED ALBUMS + feature to add stuff to created albums
// TODO implement file retrieval
func main() {
	defer os.Remove(oggcrypto.MASTERKEY_PATH)
	err := LoadDotenv()
	if err != nil {
		log.Fatal("Error loading .env file %w", err)
	}

	pgURI := os.Getenv("POSTGRES_URI")
	fmt.Println(pgURI)

	r := src.SetupRouter()

	dbl, err := src.GetDB()
	if err != nil {
		log.Fatalf("error occurred while getting the database:\n\t%v", err)
	}

	defer testing_material.FlushDB() //development mode 
	if err = user.CreateAdminUser();err != nil {
		log.Fatalf("error occurred while creating admin user:\n\t%v", err)
	}

	fmt.Print("%w", dbl)
	r.Run(":5000")
}
