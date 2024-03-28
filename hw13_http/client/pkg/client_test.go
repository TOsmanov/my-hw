package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	server "github.com/TOsmanov/my-hw/hw13_http/server/pkg"
	"github.com/stretchr/testify/assert"
)

func TestGetMessage(t *testing.T) {
	expected := "Hello world!"
	srv := httptest.NewServer(http.HandlerFunc(server.Handle))
	defer srv.Close()
	client := NewClient(srv.URL, "")
	resp, err := client.GetMessage()
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)
}

func TestPostMessage(t *testing.T) {
	expected := "Hello world!"
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		server.Handle(w, req)
	}))
	defer svr.Close()
	client1 := NewClient(svr.URL, "")
	resp, err := client1.GetMessage()
	assert.Nil(t, err)
	assert.Equal(t, expected, resp)

	firstResp := "Message is updated"
	newMessage := "Goodbye..."
	resp, err = client1.PostMessage(newMessage)
	assert.Nil(t, err)
	assert.Equal(t, firstResp, resp)

	client2 := NewClient(svr.URL, "")
	resp, err = client2.GetMessage()
	assert.Nil(t, err)
	assert.Equal(t, newMessage, resp)
}
