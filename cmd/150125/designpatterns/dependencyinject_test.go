package designpatterns

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	"weezel/meetup/internal/logger"

	"github.com/google/go-cmp/cmp"
)

type mockHTTPClient struct {
	// Make it possible to mock each test case as needed.
	// Handy in table tests.
	// Otherwise there would be a need to create a new mock for the each test case.
	mDo func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	logger.Logger.Debug().Interface("payload", req).Msg("Payload")
	return m.mDo(req)
}

func TestHandler_GetUsers(t *testing.T) {
	t.Helper()

	// This simulates a real world situation where API specification is known.
	// In this example it's known that it returns a list of users as as JSON object
	// and requires Authorization header.
	//
	//nolint:lll // This is generated payload, let it be
	msgPayloads := []string{
		`[{"name":"Alice","username":"InWonderland","age":12},{"name":"John","username":"JohnDoe","age":66},{"name":"Foo","username":"Bar","age":100}]`,
		`[{"non_existing_field": true, "name":"Alice","username":"InWonderland","age":12},{"name":"John","username":"JohnDoe","age":66},{"name":"Foo","username":"Bar","age":100}]`,
	}

	type fields struct {
		httpCli HTTPClienter
	}
	tests := []struct {
		ctx         context.Context
		fields      fields
		name        string
		expectedErr string
		want        []User
	}{
		{
			name: "Get users",
			ctx:  context.TODO(),
			fields: fields{
				httpCli: &mockHTTPClient{
					mDo: func(req *http.Request) (*http.Response, error) {
						// Check that Authorization header is incldued
						if req.Header.Get("Authorization") == "" {
							t.Error("Auth header missing")
						}

						// Check that URL is matching the specification
						if !strings.HasSuffix(req.URL.String(), "/get/users") {
							t.Errorf("Mistyped URL: %s", req.URL.String())
						}

						msg := strings.NewReader(msgPayloads[0])
						return &http.Response{
							Status:     "200 OK",
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(msg),
						}, nil
					},
				},
			},
			want: []User{
				{Name: "Alice", Username: "InWonderland", Age: 12},
				{Name: "John", Username: "JohnDoe", Age: 66},
				{Name: "Foo", Username: "Bar", Age: 100},
			},
			expectedErr: "",
		},
		{
			name: "Non existing field in payload",
			ctx:  context.TODO(),
			fields: fields{
				httpCli: &mockHTTPClient{
					mDo: func(_ *http.Request) (*http.Response, error) {
						msg := strings.NewReader(msgPayloads[1])
						return &http.Response{
							Status:     "200 OK",
							StatusCode: http.StatusOK,
							Body:       io.NopCloser(msg),
						}, nil
					},
				},
			},
			want: nil,
			// Wrapped aka chained error messages
			expectedErr: `json unmarshal: json: unknown field "non_existing_field"`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := New(WithHTTPClient(tt.fields.httpCli))
			got, err := h.GetUsers(tt.ctx)
			if (err != nil) && err.Error() != tt.expectedErr {
				t.Errorf("Handler.GetUsers() error = %q, wantErr = %q", err, tt.expectedErr)
				return
			}
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Handler.GetUsers() payload mismath:\n%s\n", diff)
			}
		})
	}
}
