package main

import (
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := loadConfig()
	if err != nil {
		t.Errorf("NO CONFIG FILE")
	}
}
func TestServer(t *testing.T) {
	err := startServer("test.html")
	if err == nil {
		t.Errorf("SERVER STARTED WITH INVALID FILE")
	}
}
