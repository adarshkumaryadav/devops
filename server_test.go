package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHelloWorldShouldPass(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(Greet))
	// close at the last at any cost
	defer testServer.Close()

	testClient := testServer.Client()
	response, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}
	/*
		 	if response.StatusCode != 200 {
				t.Error("expecting 200 status code")
			}
	*/
	body, err := io.ReadAll(response.Body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, response.StatusCode, 200)
	assert.Equal(t, "hello world", string(body))

}

func TestHelloWorldShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(Greet))

	testClient := testServer.Client()
	body := strings.NewReader("body...")
	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, resp.StatusCode, 405)

}
func TestHealth(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))
	testClient := testServer.Client()

	resp, err := testClient.Get(testServer.URL)
	if err != nil {
		t.Error(err.Error())
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, resp.StatusCode, 200)
	assert.Equal(t, string(body), "hello, I am alive...")
}

func TestHealthShouldFail(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handleHealth))

	testClient := testServer.Client()
	body := strings.NewReader("body...")
	resp, err := testClient.Post(testServer.URL, "application/json", body)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, resp.StatusCode, 405)

}
