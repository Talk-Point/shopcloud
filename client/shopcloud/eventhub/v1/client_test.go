package client

import (
	"testing"
)

func TestClient(t *testing.T) {
	client := NewCloudClient("http://localhost:8080", "token")
	if client == nil {
		t.Fatal("client is nil")
	}
}
