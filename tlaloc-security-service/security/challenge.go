package security

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"github.com/labstack/gommon/log"
	"time"
)

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
	// 1. Construir el string original: password + challenge + nonce + timestamp

	log.Info("pwd" + password)
	originalString := fmt.Sprintf("%s%s%s%d", password, challenge, nonce, timestamp)

	// 2. Primer hash: SHA256(password + challenge + nonce + timestamp)
	firstHash := sha256.Sum256([]byte(originalString))
	firstHashStr := hex.EncodeToString(firstHash[:])

	// 3. Segundo hash: SHA256(primer_hash) - Double hashing
	secondHash := sha256.Sum256([]byte(firstHashStr))
	expectedHash := hex.EncodeToString(secondHash[:])

	// 4. Comparación constante en tiempo
	return subtle.ConstantTimeCompare([]byte(clientHash), []byte(expectedHash)) == 1, nil
}

// IsChallengeValid - Verifica si el challenge es válido y no ha expirado
func IsChallengeValid(challenge *ChallengeResponse) bool {
	return challenge != nil &&
		GetCurrentTime().Before(challenge.ExpiresAt) &&
		len(challenge.Challenge) == 64 &&
		len(challenge.Nonce) == 32
}
