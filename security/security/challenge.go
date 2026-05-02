package security

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomString(length int) (string, error) {
	result := make([]byte, length)
	for i := range result {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}

func GetCurrentTime() time.Time {
	return time.Now()
}

// ChallengeResponse - Estructura para el challenge
type ChallengeResponse struct {
	Challenge string    `json:"challenge"`
	ExpiresAt time.Time `json:"expires_at"`
	Timestamp time.Time `json:"timestamp"`
	Nonce     string    `json:"nonce"`
}

// GenerateChallenge - Genera un challenge seguro
func GenerateChallenge() (*ChallengeResponse, error) {
	// 1. Nonce aleatorio para prevenir replay attacks
	nonce, err := GenerateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate nonce: %w", err)
	}

	// 2. Challenge aleatorio
	challenge, err := GenerateRandomString(64)
	if err != nil {
		return nil, fmt.Errorf("failed to generate challenge: %w", err)
	}

	// 3. El challenge expira en 5 minutos
	expiresAt := GetCurrentTime().Add(5 * time.Minute)

	return &ChallengeResponse{
		Challenge: challenge,
		ExpiresAt: expiresAt,
		Timestamp: GetCurrentTime(),
		Nonce:     nonce,
	}, nil
}

// VerifyChallengeHash - Verifica el hash del challenge
func VerifyChallengeHash(password, challenge, nonce, clientHash string, timestamp int64) (bool, error) {
	fmt.Println(password)

	passwordHash := sha256.Sum256([]byte(password))
	passwordHashStr := hex.EncodeToString(passwordHash[:])
	fmt.Println(passwordHashStr)

	originalString := fmt.Sprintf("%s%s%s%d", passwordHashStr, challenge, nonce, timestamp)

	firstHash := sha256.Sum256([]byte(originalString))
	firstHashStr := hex.EncodeToString(firstHash[:])

	secondHash := sha256.Sum256([]byte(firstHashStr))
	expectedHash := hex.EncodeToString(secondHash[:])
	fmt.Print(expectedHash)

	return subtle.ConstantTimeCompare([]byte(clientHash), []byte(expectedHash)) == 1, nil
}

// IsChallengeValid - Verifica si el challenge es válido y no ha expirado
func IsChallengeValid(challenge *ChallengeResponse) bool {
	return challenge != nil &&
		GetCurrentTime().Before(challenge.ExpiresAt) &&
		len(challenge.Challenge) == 64 &&
		len(challenge.Nonce) == 32
}
