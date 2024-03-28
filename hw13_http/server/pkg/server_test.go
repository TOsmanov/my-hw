package server

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleGet(t *testing.T) {
	expected := "Hello world!"
	req := httptest.NewRequest(http.MethodGet, "http://localhost:10001", nil)
	w := httptest.NewRecorder()
	Handle(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, expected, string(body))
}

func TestHandlePost(t *testing.T) {
	firstResp := "Message is updated"
	expected := "Goodbye..."
	body := strings.NewReader(expected)
	req := httptest.NewRequest(http.MethodPost, "http://localhost:10001", body)
	w := httptest.NewRecorder()
	Handle(w, req)

	resp := w.Result()
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, firstResp, string(respBody))

	req = httptest.NewRequest(http.MethodGet, "http://localhost:10001", nil)
	w = httptest.NewRecorder()
	Handle(w, req)

	resp = w.Result()
	defer resp.Body.Close()
	respBody, err = io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, expected, string(respBody))
}
