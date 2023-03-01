package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestDownloadUrl(t *testing.T) {
	// Test server setup
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test data"))
	}))
	defer ts.Close()

	// Test case data
	testURL := ts.URL
	expectedContent := []byte("test data")

	// Test function call
	err := DownloadUrl(testURL)
	if err != nil {
		t.Errorf("DownloadUrl failed with error: %v", err)
	}

	// Check if file was created and its content
	filename := filepath.Join(".", "localhost", "index.html")
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("DownloadUrl failed to create file: %v", err)
	}

	// Compare expected and actual results
	if !bytes.Equal(content, expectedContent) {
		t.Errorf("DownloadUrl produced incorrect file content, got %s, want %s", content, expectedContent)
	}
}
