package general

import (
	"bytes"
	"encoding/json"
	"fmt"

	"weezel/meetup/internal/logger"
)

type MyStruct struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func enforcedJSONdecode() {
	input := `{
		"name": "Test user",
		"age": 100,
		"email": "test.user@example.com",
		"extra": "unexpected"
	}`

	decoder := json.NewDecoder(bytes.NewReader([]byte(input)))
	decoder.DisallowUnknownFields()

	var data MyStruct
	err := decoder.Decode(&data)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to decode JSON")
		return
	}

	fmt.Printf("Decoded JSON: %#v\n", data)
}

func relaxedJSONdecode() {
	input := `{
		"name": "Test user",
		"age": 100,
		"email": "test.user@example.com",
		"extra": "unexpected"
	}`

	var data MyStruct
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Failed to decode JSON")
		return
	}

	fmt.Printf("Decoded JSON: %#v\n", data)
}
