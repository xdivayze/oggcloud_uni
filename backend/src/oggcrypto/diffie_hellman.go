package oggcrypto

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"fmt"

	"golang.org/x/crypto/hkdf"
)

func GenerateECDHPair() (*ecdh.PrivateKey, *ecdh.PublicKey, error) {
	privateKey, err := ecdh.P256().GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, fmt.Errorf("error generating ecdh private key:\n\t%w", err)
	}
	return privateKey, privateKey.PublicKey(), nil

}

func DeriveSharedSecret(privateKey *ecdh.PrivateKey, publicKey *ecdh.PublicKey, saltopt interface{}) ([]byte, []byte, error) {
	shared, err := privateKey.ECDH(publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error deriving shared secret:\n\t%w", err)
	}
	var salt []byte
	if saltopt == nil {
		salt = make([]byte, SALT_LENGTH)
		_, err = rand.Read(salt)
		if err != nil {
			return nil, nil, fmt.Errorf("error reading from random buffer:\n\t%w", err)
		}
	} else {
		var ok bool
		salt, ok = saltopt.([]byte)
		if !ok {
			return nil, nil, fmt.Errorf("given salt cannot be casted to type []byte")
		}
	}

	hkdf := hkdf.New(sha256.New, shared, salt, nil)
	derivedKey := make([]byte, AES_KEY_SIZE)
	_, err = hkdf.Read(derivedKey)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading from hkdf buffer:\n\t%w", err)
	}
	return derivedKey, salt, nil
}

func ReadFromPEM(pemBlock string) (*ecdh.PublicKey, error) {
	pemEncoded, err := hex.DecodeString(pemBlock)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 encoded pem:\n\t%w", err)
	}
	pubBytes, _ := pem.Decode(pemEncoded)
	pubKey, err := x509.ParsePKIXPublicKey(pubBytes.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error parsing public key:\n\t%w", err)
	}
	pubKeyConverted, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("type casting failed")
	}
	pubkeyNIST, err := pubKeyConverted.ECDH()
	if err != nil {
		return nil, fmt.Errorf("error converting from ecdsa to ecdh:\n\t%w", err)
	}
	return pubkeyNIST, nil
}

func EncodePublicKeyToPEM(publicKey *ecdh.PublicKey) (string, error) {
	pubBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", fmt.Errorf("error serializing public key:\n\t%w", err)
	}
	block := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubBytes,
	}

	pemEncoded := pem.EncodeToMemory(block)
	hexEncoded := hex.EncodeToString(pemEncoded)
	return hexEncoded, nil
}
