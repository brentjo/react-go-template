package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	setupServer()
	code := m.Run()
	os.Exit(code)
}

func TestStaticFilesHandler(t *testing.T) {
	testServer := httptest.NewServer(nil)
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/index.html")
	if err != nil {
		t.Fatalf("Failed to make request to /index.html: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	expectedContentType := "text/html"
	if !strings.Contains(resp.Header.Get("Content-Type"), expectedContentType) {
		t.Errorf("Expected Content-Type to include %s, got %s", expectedContentType, resp.Header.Get("Content-Type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expectedContentSnippet := `<div id="root"></div>`
	if !strings.Contains(string(body), expectedContentSnippet) {
		t.Errorf("Expected response body to contain %q", expectedContentSnippet)
	}
}

func TestAllowedSPAPaths(t *testing.T) {
	testServer := httptest.NewServer(nil)
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/time")
	if err != nil {
		t.Fatalf("Failed to make request to /time: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	expectedContentType := "text/html"
	if !strings.Contains(resp.Header.Get("Content-Type"), expectedContentType) {
		t.Errorf("Expected Content-Type to include %s, got %s", expectedContentType, resp.Header.Get("Content-Type"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	expectedContentSnippet := `<div id="root"></div>`
	if !strings.Contains(string(body), expectedContentSnippet) {
		t.Errorf("Expected response body to contain %q", expectedContentSnippet)
	}
}

func TestTimeHandler(t *testing.T) {
	testServer := httptest.NewServer(nil)
	defer testServer.Close()

	resp, err := http.Get(testServer.URL + "/api/time")
	if err != nil {
		t.Fatalf("Failed to make request to /api/time: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var data struct {
		Time string `json:"time"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if data.Time == "" {
		t.Errorf("Expected 'time' key to be present and not empty")
	}
}
