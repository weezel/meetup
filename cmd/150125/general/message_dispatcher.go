package general

import (
	"encoding/json"
	"errors"
	"fmt"

	"weezel/meetup/internal/logger"
)

var ErrUnkownType = errors.New("unknown type")

type Email struct {
	From    string   `json:"from,omitempty"`
	To      string   `json:"to,omitempty"`
	Subject string   `json:"subject,omitempty"`
	Cc      []string `json:"cc,omitempty"`
	Bcc     []string `json:"bcc,omitempty"`
	Body    []byte   `json:"body,omitempty"`
}

type SMS struct {
	From       string `json:"from,omitempty"`
	To         string `json:"to,omitempty"`
	RcptNumber string `json:"rcpt_number,omitempty"`
	Text       string `json:"text,omitempty"`
}

// CommonMessage struct contains unique fields from the Email and from the SMS
// structs so that it's possible to determine which message the payload really is
type CommonMessage struct {
	RcptNumber string `json:"rcpt_number,omitempty"`
	Body       string `json:"body,omitempty"`
}

func MessageDispatcher(payload []byte) error {
	common := CommonMessage{}
	if err := json.Unmarshal(payload, &common); err != nil {
		return fmt.Errorf("unmarshal: %w", err)
	}

	switch {
	case common.RcptNumber != "":
		sms := SMS{}
		if err := json.Unmarshal(payload, &sms); err != nil {
			logger.Logger.Error().Err(err).Msg("Failed to unmarshal to SMS")
		}
		logger.Logger.Info().Interface("sms", sms).Msg("Got an SMS payload")
	case len(common.Body) > 0:
		email := Email{}
		if err := json.Unmarshal(payload, &email); err != nil {
			logger.Logger.Error().Err(err).Msg("Failed to unmarshal to Email")
		}
		logger.Logger.Info().Interface("email", email).Msg("Got an Email payload")
	default:
		logger.Logger.Warn().Bytes("payload", payload).Msg("Unknown payload")
		return ErrUnkownType
	}

	return nil
}
