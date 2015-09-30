package discovery_test

import (
	"log"
	"os"
	"testing"

	"github.com/Clever/discovery-go"
)

func insertPairsToEnv(pairs map[string]string) {
	for name, val := range pairs {
		err := os.Setenv(name, val)
		if err != nil {
			log.Fatalf("Failed to set env variable, %s", err)
		}
	}
}

func TestMain(m *testing.M) {
	insertPairsToEnv(map[string]string{
		"SERVICE_REDIS_TCP_URL":         "tcp://redis.com:6379",
		"SERVICE_GOOGLE_API_URL":        "https://api.google.com:80",
		"SERVICE_LONG_APP_NAME_API_URL": "http://long-app-name:80",
	})

	os.Exit(m.Run())
}

func TestTCPDiscovery(t *testing.T) {
	expected := "tcp://redis.com:6379"
	expectedScheme := "tcp"

	url, err := discovery.URLString("redis", "tcp")
	if err != nil {
		t.Fatalf("Unexpected error, %s", err)
	} else if url != expected {
		t.Fatalf("Unexpected result, expected: %s, received: %s", expected, url)
	}

	u, err := discovery.URL("redis", "tcp")
	if err != nil {
		t.Fatalf("Unexpected error, %s", err)
	} else if u.String() != expected {
		t.Fatalf("unexpected result, expected: %s, received: %s", expected, u.String())
	} else if u.Scheme != expectedScheme {
		t.Fatalf("unexpected result, expected: %s, received: %s", expectedScheme, u.Scheme)
	}
}

func TestHTTPSDiscovery(t *testing.T) {
	expected := "https://api.google.com:80"

	url, err := discovery.URLString("google", "api")
	if err != nil {
		t.Fatalf("Unexpected error, %s", err)
	} else if url != expected {
		t.Fatalf("Unexpected result, expected: %s, received: %s", expected, url)
	}
}

func TestErrorOnFailure(t *testing.T) {
	_, err := discovery.URLString("break", "api")
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestLongArbitraryNameWithDashes(t *testing.T) {
	_, err := discovery.URLString("long-app-name", "api")
	if err != nil {
		t.Fatalf("Unexpected error with app name w/ dashes, %s", err)
	}
}
