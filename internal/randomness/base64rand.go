package randomness

import (
	"crypto/rand"
	"encoding/base64"

	"weezel/meetup/internal/logger"
)

const length = 12

func GenerateRandomCode() string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		logger.Logger.Fatal().Err(err).Msg("Randomness has failed us")
	}

	randomString := base64.URLEncoding.EncodeToString(bytes)
	return randomString[:length]
}
