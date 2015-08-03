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
	})

	os.Exit(m.Run())
}

func TestTCPDiscovery(t *testing.T) {
	expected := "tcp://redis.com:6379"

	url := discovery.Discover("redis", "tcp")
	if url != expected {
		t.Fatalf("Unexpected result, expected: %s, receieved: %s", expected, url)
	}
}

func TestHTTPSDiscovery(t *testing.T) {
	expected := "https://api.google.com:80"

	url := discovery.Discover("google", "api")
	if url != expected {
		t.Fatalf("Unexpected result, expected: %s, receieved: %s", expected, url)
	}
}

func TestPanicOnFailure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Discover() should panic with missing environment variables")
		}
	}()
	_ = discovery.Discover("break", "api")
}
