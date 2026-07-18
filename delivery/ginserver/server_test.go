package ginserver

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestPingContract(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.GET("/ping", pingHandler)

	request := httptest.NewRequest(http.MethodGet, "/ping", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, response.Code)
	}

	var body struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if body.Message != "pong" {
		t.Fatalf("expected pong message, got %q", body.Message)
	}
}

func TestLocalCORSConfig(t *testing.T) {
	config := localCORSConfig()
	for _, origin := range []string{"http://localhost:8080", "http://127.0.0.1:3000", "https://[::1]:8443"} {
		if !config.AllowOriginFunc(origin) {
			t.Errorf("expected local origin %q to be allowed", origin)
		}
	}
	for _, origin := range []string{"https://example.com", "https://192.168.1.10", "file:///tmp/index.html", "not a URL"} {
		if config.AllowOriginFunc(origin) {
			t.Errorf("expected non-local origin %q to be rejected", origin)
		}
	}
}
