package api

import (
	"testing"
	"net/http/httptest"
	"net/http"
)


func init(){
}

func TestHealthCheckHandler(t *testing.T) {
	server := httptest.NewServer(Handlers())
	defer server.Close()
	resp, err := http.Get(server.URL + "/healthcheck")

	if err != nil {
		t.Fatal("ERROR: " + err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Received non 200 response %d", resp.StatusCode)
	}
}

