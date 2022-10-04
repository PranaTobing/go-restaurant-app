package user

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

func (ur *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	argonHash := argon2.IDKey(
		[]byte(password),
		salt,
		ur.time,
		ur.memory,
		ur.parallelism,
		ur.keyLen,
	)

	b64Hash := ur.encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, ur.memory, ur.time, ur.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (ur *userRepo) encrypt(text []byte) string {
	nonce := make([]byte, ur.gcm.NonceSize())

	ciphertext := ur.gcm.Seal(nonce, nonce, text, nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

func (ur *userRepo) decrypt(ciphertext string) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return nil, err
	}
	if len(decoded) < ur.gcm.NonceSize() {
		return nil, err
	}

	return ur.gcm.Open(nil,
		decoded[:ur.gcm.NonceSize()],
		decoded[ur.gcm.NonceSize():],
		nil,
	)
}
