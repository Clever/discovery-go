package discovery

import (
	"fmt"
	"net/url"
	"os"
	"strings"
)

const (
	templateVar = "SERVICE_%s_%s_URL"
)

// URLString finds the specified URL for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_URL.
func URLString(service, name string) (string, error) {
	// build string
	envKey := fmt.Sprintf(templateVar, service, name)

	// standardize to uppercase & underscores
	envKey = strings.Replace(strings.ToUpper(envKey), "-", "_", -1)

	val := os.Getenv(envKey)
	if val == "" {
		return "", fmt.Errorf("Missing env var, %s", envKey)
	}
	return val, nil
}

// URL finds the specified URL for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_URL.
func URL(service, name string) (*url.URL, error) {
	// get string
	s, err := URLString(service, name)
	if err != nil {
		return nil, err
	}

	return url.Parse(s)
}
