package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

	ht := httptest.NewServer(http.HandlerFunc(handler))


	resp, err := http.Get(ht.URL + "/")
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, []byte("Hello\n"), body, "Unexpected response")
}
