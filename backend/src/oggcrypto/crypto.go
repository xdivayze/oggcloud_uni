package oggcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func init() {
	key, err := GenerateAESKey()
	if err != nil {
		log.Fatalf("error generating aes keys:\n\t%v", err)
	}
	keyHex := hex.EncodeToString(key)
	file, err := os.Create(MASTERKEY_PATH)
	if err != nil {
		log.Fatalf("error creating file at path %s :\n\t%v", MASTERKEY_PATH, err)
	}
	defer file.Close()
	_, err = file.WriteString(keyHex)
	if err != nil {
		log.Fatalf("failed to write to file:\n\t%v", err)
	}
}

func DecryptAES(ciphertext []byte, key []byte) ([]byte, error) {
	if len(ciphertext) < 12 {
		return nil, fmt.Errorf("ciphertext too short to contain a nonce")
	}
	nonce := ciphertext[:NONCE_LENGTH]
	encData := ciphertext[NONCE_LENGTH:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not create AES cipher: \n\t%v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("could not create GCM mode: \n\t%v", err)
	}

	plainText, err := aesGCM.Open(nil, nonce, encData, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed:\n\t%v", err)
	}
	return plainText, nil

}

func GenerateAESKey() ([]byte, error) {
	key := make([]byte, AES_KEY_SIZE)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		return nil, fmt.Errorf("error reading random to buffer:\n\t%w", err)
	}
	return key, nil
}

func EncryptAES(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("could not create aes cipher:\n\t %w", err)
	}
	nonce := make([]byte, NONCE_LENGTH)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("could not generate nonce:\n\t%w", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("couldnt generate gcm mode:\n\t%w", err)
	}

	cipherText := aesGCM.Seal(nil, nonce, data, nil)
	return append(nonce, cipherText...), nil

}
