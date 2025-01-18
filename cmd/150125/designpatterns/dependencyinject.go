package designpatterns

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

var serviceURL, _ = url.Parse("http://localhost/such/nice/such/wow")

type HTTPClienter interface {
	// Method signature of http.Client.Do()
	Do(req *http.Request) (*http.Response, error)
}

type Option func(*Handler)

func WithHTTPClient(client HTTPClienter) Option {
	return func(h *Handler) {
		h.httpCli = client
	}
}

func WithAuth(auth string) Option {
	return func(h *Handler) {
		h.authHeader = auth
	}
}

type Handler struct {
	httpCli    HTTPClienter
	authHeader string
}

func New(opts ...Option) *Handler {
	h := &Handler{
		httpCli:    &http.Client{Timeout: time.Second * 30},
		authHeader: "demo",
	}

	// Override defaults
	for _, opt := range opts {
		opt(h)
	}

	return h
}

// User	defines user JSON object as per demo specification
type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Age      int    `json:"age,omitempty"`
}

// GetUsers imitates a call to the third party service which returns a list of users.
func (h *Handler) GetUsers(ctx context.Context) ([]User, error) {
	usersURL := serviceURL.JoinPath("get/users")
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, usersURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("url request: %w", err)
	}
	req.Header.Add("Authorization", h.authHeader)

	resp, err := h.httpCli.Do(req)
	if err != nil {
		return nil, fmt.Errorf("get user list: %w", err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	payload, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields()
	var users []User
	if err = decoder.Decode(&users); err != nil {
		return nil, fmt.Errorf("json unmarshal: %w", err)
	}

	return users, nil
}

// Usage in
func exampleUsage() { //nolint:unused // Demo
	ctx := context.Background()
	// Create a new Handler by calling the constructor which uses dependency injection
	handler := New(
		WithHTTPClient(http.DefaultClient),
		WithAuth("Bearer aabbcc"),
	)
	// Since this is in the same package space, it's possible to call httpCli field directly.
	// Outside this package space the only available method is `GetUser()`.
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "http://example.com", nil)
	//nolint:errcheck,bodyclose // Error checking omitted for previty. Also ignore body close in demo
	handler.httpCli.Do(req)
}
