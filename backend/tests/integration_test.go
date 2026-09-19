package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	Tenant   string `json:"tenant"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type LogItem struct {
	ID        uint   `json:"id"`
	Tenant    string `json:"tenant"`
	Source    string `json:"source"`
	EventType string `json:"event_type"`
	User      string `json:"user"`
	SrcIP     string `json:"src_ip"`
}

type LogResponse struct {
	Count int       `json:"count"`
	Logs  []LogItem `json:"logs"`
}

func getAPIBaseURL() string {
	if value := os.Getenv("API_BASE_URL"); value != "" {
		return value
	}

	return "http://localhost/api"
}

func login(
	t *testing.T,
	username string,
	password string,
) LoginResponse {
	t.Helper()

	requestBody := LoginRequest{
		Username: username,
		Password: password,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatalf("failed to encode login request: %v", err)
	}

	response, err := http.Post(
		getAPIBaseURL()+"/auth/login",
		"application/json",
		bytes.NewBuffer(body),
	)
	if err != nil {
		t.Fatalf("login request failed: %v", err)
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatalf("failed to read login response: %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected login status 200, got %d: %s",
			response.StatusCode,
			string(responseBody),
		)
	}

	var loginResponse LoginResponse

	if err := json.Unmarshal(
		responseBody,
		&loginResponse,
	); err != nil {
		t.Fatalf(
			"failed to decode login response: %v",
			err,
		)
	}

	return loginResponse
}

func authenticatedGET(
	t *testing.T,
	url string,
	token string,
) *http.Response {
	t.Helper()

	request, err := http.NewRequest(
		http.MethodGet,
		url,
		nil,
	)
	if err != nil {
		t.Fatalf(
			"failed to create request: %v",
			err,
		)
	}

	request.Header.Set(
		"Authorization",
		"Bearer "+token,
	)

	response, err := http.DefaultClient.Do(
		request,
	)
	if err != nil {
		t.Fatalf(
			"request failed: %v",
			err,
		)
	}

	return response
}

func TestHealthEndpoint(t *testing.T) {
	response, err := http.Get(
		getAPIBaseURL() + "/health",
	)
	if err != nil {
		t.Fatalf(
			"health request failed: %v",
			err,
		)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf(
			"expected status 200, got %d",
			response.StatusCode,
		)
	}

	var result map[string]string

	if err := json.NewDecoder(
		response.Body,
	).Decode(&result); err != nil {
		t.Fatalf(
			"failed to decode health response: %v",
			err,
		)
	}

	if result["status"] != "ok" {
		t.Errorf(
			"expected status ok, got %q",
			result["status"],
		)
	}
}

func TestAdminLogin(t *testing.T) {
	response := login(
		t,
		"admin",
		"admin123",
	)

	if response.Token == "" {
		t.Fatal(
			"expected JWT token, got empty token",
		)
	}

	if response.User.Username != "admin" {
		t.Errorf(
			"expected username admin, got %q",
			response.User.Username,
		)
	}

	if response.User.Role != "admin" {
		t.Errorf(
			"expected role admin, got %q",
			response.User.Role,
		)
	}
}

func TestViewerTenantIsolation(t *testing.T) {
	loginResponse := login(
		t,
		"viewerA",
		"viewer123",
	)

	// Viewer intentionally asks for demoB.
	// Backend must ignore this and force viewer's own tenant.
	url := fmt.Sprintf(
		"%s/logs?tenant=demoB",
		getAPIBaseURL(),
	)

	response := authenticatedGET(
		t,
		url,
		loginResponse.Token,
	)
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(
			response.Body,
		)

		t.Fatalf(
			"expected status 200, got %d: %s",
			response.StatusCode,
			string(body),
		)
	}

	var result LogResponse

	if err := json.NewDecoder(
		response.Body,
	).Decode(&result); err != nil {
		t.Fatalf(
			"failed to decode logs response: %v",
			err,
		)
	}

	if len(result.Logs) == 0 {
		t.Fatal(
			"expected demoA logs for viewerA, got no logs",
		)
	}

	for _, logItem := range result.Logs {
		if logItem.Tenant != "demoA" {
			t.Errorf(
				"tenant isolation failed: expected demoA, got %q for log ID %d",
				logItem.Tenant,
				logItem.ID,
			)
		}
	}
}
