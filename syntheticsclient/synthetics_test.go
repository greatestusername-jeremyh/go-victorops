package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var (
	// mux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// server is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

func TestConfigurableClient(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)
	args := ClientArgs{
		timeoutSeconds: 30,
		publicBaseUrl:  testServer.URL,
	}

	testConfigurableClient := NewConfigurableClient("snakedonut", args)
	log.Printf("Client instantiated: %s", testServer.URL)
	if testConfigurableClient.GetHTTPClient() == nil {
		t.Errorf("http client is nil")
	}
	if testConfigurableClient.apiKey != "snakedonut" {
		t.Errorf("apiKey was not correctly passed")
	}
}

func TestConfigurableClientTimeout(t *testing.T) {
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	testMux.HandleFunc("/v2/checks/12", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
	})


	testConfigurableClient := NewConfigurableClient("apiKey", ClientArgs{
		timeoutSeconds: 1,
		publicBaseUrl:  testServer.URL,
	})
	log.Printf("Client instantiated: %s", testServer.URL)
	_, _, err := testConfigurableClient.GetCheck(12)
	if !strings.Contains(err.Error(), "context deadline exceeded (Client.Timeout exceeded while awaiting headers)") {
		t.Errorf("expected to see timeout error, but saw: %s", err.Error())
	}
}
