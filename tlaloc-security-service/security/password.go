package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/argon2"
)

func HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("failed to generate salt: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	return fmt.Sprintf(format, argon2.Version, 64*1024, 1, 4, b64Salt, b64Hash), nil
}

func VerifyPassword(password, hash string) (bool, error) {
	var version, memory, time, parallelism uint32
	var salt, hashBytes []byte

	_, err := fmt.Sscanf(hash, "$argon2id$v=%d$m=%d,t=%d,p=%d$%x$%x",
		&version, &memory, &time, &parallelism, &salt, &hashBytes)
	if err != nil {
		return false, fmt.Errorf("invalid hash format")
	}

	computedHash := argon2.IDKey([]byte(password), salt, time, memory, uint8(parallelism), uint32(len(hashBytes)))

	return subtle.ConstantTimeCompare(hashBytes, computedHash) == 1, nil
}

func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func GetCurrentTime() time.Time {
	return time.Now().UTC()
}

func AddDuration(t time.Time, duration time.Duration) time.Time {
	return t.Add(duration)
}
