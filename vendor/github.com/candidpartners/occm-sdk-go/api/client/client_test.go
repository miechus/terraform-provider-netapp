// Package implements helper client functionality.
// TODO: add more comprehensive test cases
package client

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

  "github.com/candidpartners/occm-sdk-go/util"
	"github.com/stretchr/testify/assert"
)

var (
	mux *http.ServeMux
	client *Client
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
  context := &Context{
    Host: "1.2.3.4",
  }
  client, _ = New(context)
  client.apiUrl = server.URL
}

func teardown() {
	server.Close()
}

func TestInvoke(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/some/random/endpoint", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "PUT", r.Method, "Expected method 'PUT', got %s", r.Method)

    body, err := util.FromJSONStream(r.Body)
    if err != nil {
      t.Fatal(err, "Error parsing body")
    }

    want := map[string]interface{}{
      "boolInput": false,
      "stringOutput": "in123",
    }

    if assert.NoError(t, err) {
  		assert.Equal(t, body, want)
  	}

    fmt.Fprintf(w, `{
  "boolOutput": true,
  "stringOutput": "out456"
  }`)
	})

	response, err := client.Invoke("PUT", "/some/random/endpoint",
    map[string]string{}{
      "qs1": "qs789",
      "qs2": false,
    },
    map[string]interface{}{
      "stringOutput": "in123",
      "boolInput": false,
  });

  body, err := util.FromJSON(response)
  if err != nil {
    t.Fatal(err, "Error parsing response")
  }

  want := map[string]interface{}{
    "stringOutput": "out456",
    "boolOutput": true,
  }

  if assert.NoError(t, err) {
    assert.Equal(t, body, want)
  }

	assert.NoError(t, err)
}
