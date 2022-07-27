package envar

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

func HashID(tgID int64) (string, error) {

	salt := []byte(os.Getenv("salt"))

	IDBytes := []byte(fmt.Sprintf("%v", tgID))

	sha512Hasher := sha256.New()

	IDBytes = append(IDBytes, salt...)

	_, err := sha512Hasher.Write(IDBytes)
	if err != nil {
		return "", err
	}

	hashedIDBytes := sha512Hasher.Sum(nil)

	var base64EncodedPasswordHash = base64.URLEncoding.EncodeToString(hashedIDBytes)

	return base64EncodedPasswordHash, nil

}
