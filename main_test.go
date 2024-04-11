package main

import (
	"os"
	"testing"
)

func TestGetEnv(t *testing.T) {
	if err := os.Setenv("HTTP_PORT", "9086"); err != nil {
		t.Fatal(err)
	}

	v := getEnv("HTTP_PORT", "8080")
	if v != "9086" {
		t.Errorf("value is not valid,expected : %s, got : %s", "9086", v)
	}
}
