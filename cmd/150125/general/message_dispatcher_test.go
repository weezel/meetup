package general

import (
	"encoding/json"
	"testing"

	"weezel/meetup/internal/logger"
)

func emailPayload() []byte {
	email := Email{
		From:    "user@example.com",
		To:      "foo@bar.com",
		Subject: "Testing",
		Body:    []byte("Hello from Email!"),
	}
	buf, err := json.MarshalIndent(email, "", "  ")
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to marshal email")
	}
	return buf
}

func smsPayload() []byte {
	sms := SMS{
		From:       "Alice",
		To:         "Bob",
		RcptNumber: "12345",
		Text:       "Hello from SMS!",
	}
	buf, err := json.MarshalIndent(sms, "", "  ")
	if err != nil {
		logger.Logger.Fatal().Err(err).Msg("Failed to marshal SMS")
	}
	return buf
}

func TestMessageDispatcher(t *testing.T) {
	t.Helper()

	type args struct {
		payload []byte
	}
	tests := []struct {
		name    string
		wantErr string
		args    args
	}{
		{
			name: "Email payload",
			args: args{
				payload: emailPayload(),
			},
			wantErr: "",
		},
		{
			name: "SMS payload",
			args: args{
				payload: smsPayload(),
			},
			wantErr: "",
		},
		{
			name: "Broken JSON",
			args: args{
				payload: []byte("crash"),
			},
			wantErr: "unmarshal: invalid character 'c' looking for beginning of value",
		},
		{
			name: "Unknown payload",
			args: args{
				payload: []byte("{}"), // Minimal JSON
			},
			wantErr: ErrUnkownType.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := MessageDispatcher(tt.args.payload)
			if err != nil && tt.wantErr != err.Error() {
				t.Errorf("MessageDispatcher() got error = %q, wantErr = %q", err, tt.wantErr)
			}
		})
	}
}
