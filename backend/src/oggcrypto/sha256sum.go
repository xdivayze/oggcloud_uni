package oggcrypto

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

func CalculateSHA256sum(w io.Reader) (string, error) {
	hasher := sha256.New()
	buf := make([]byte, 1024*4)
	if _, err := io.CopyBuffer(hasher, w, buf); err != nil {
		return "", fmt.Errorf("error occured while copying through buffer:\n\t%w", err)
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}