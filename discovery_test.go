package discovery_test

import (
	"log"
	"os"
	"testing"

	"github.com/Clever/discovery-go"
)

func insertPairs(pairs map[string]string) {
	for name, val := range pairs {
		err := os.Setenv(name, val)
		if err != nil {
			log.Fatalf("Failed to set env variable, %s", err)
		}
	}
}

func TestMain(m *testing.M) {
	insertPairs(map[string]string{
		"SERVICE_REDIS_TCP_PROTO": "tcp",
		"SERVICE_REDIS_TCP_HOST":  "redis.com",
		"SERVICE_REDIS_TCP_PORT":  "6379",

		"SERVICE_GOOGLE_API_PROTO": "https",
		"SERVICE_GOOGLE_API_HOST":  "api.google.com",
		"SERVICE_GOOGLE_API_PORT":  "80",

		"SERVICE_BREAK_API_HOST": "missing.proto",
		"SERVICE_BREAK_API_PORT": "5000",

		"SERVICE_LONG_APP_NAME_API_HOST": "arbitrary",
	})

	os.Exit(m.Run())
}

func TestTCPDiscovery(t *testing.T) {
	expected := "tcp://redis.com:6379"

	url, err := discovery.URL("redis", "tcp")
	if err != nil {
		t.Fatalf("Unexpected error, %s", err)
	} else if url != expected {
		t.Fatalf("Unexpected result, expected: %s, receieved: %s", expected, url)
	}
}

func TestHTTPSDiscovery(t *testing.T) {
	expected := "https://api.google.com:80"

	url, err := discovery.URL("google", "api")
	if err != nil {
		t.Fatalf("Unexpected error, %s", err)
	} else if url != expected {
		t.Fatalf("Unexpected result, expected: %s, receieved: %s", expected, url)
	}
}

func TestErrorOnFailure(t *testing.T) {
	_, err := discovery.URL("break", "api")
	if err == nil {
		t.Fatalf("Expected error")
	}
}

func TestLongArbitraryNameWithDashesa(t *testing.T) {
	_, err := discovery.Host("long-app-name", "api")
	if err != nil {
		t.Fatalf("Unexpected error with app name w/ dashes, %s", err)
	}
}
