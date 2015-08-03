package discovery

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"strings"

	kv "gopkg.in/clever/kayvee-go.v2"
)

// m is a convenience type for using kv.
type m map[string]interface{}

const (
	templateVar = "SERVICE_%s_%s_%%s"
)

func getVar(envVar string) (string, error) {
	envVar = strings.ToUpper(envVar)
	val := os.Getenv(envVar)
	if val == "" {
		return "", errors.New(kv.FormatLog("discovery-go", kv.Error, "missing env var", m{
			"var": envVar,
		}))
	}
	return val, nil
}

// DiscoverURL finds the specified URL for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_{PROTO,HOST,PORT}.
func DiscoverURL(service, name string) (string, error) {
	proto, err := DiscoverProto(service, name)
	if err != nil {
		return "", err
	}
	host, err := DiscoverHost(service, name)
	if err != nil {
		return "", err
	}
	port, err := DiscoverPort(service, name)
	if err != nil {
		return "", err
	}

	u := url.URL{
		Scheme: proto,
		Host:   fmt.Sprintf("%s:%s", host, port),
	}
	return u.String(), nil
}

// DiscoverProto finds the specified protocol for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_PROTO.
func DiscoverProto(service, name string) (string, error) {
	template := fmt.Sprintf(templateVar, service, name)

	return getVar(fmt.Sprintf(template, "PROTO"))
}

// DiscoverHost finds the specified host for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_HOST.
func DiscoverHost(service, name string) (string, error) {
	template := fmt.Sprintf(templateVar, service, name)

	return getVar(fmt.Sprintf(template, "HOST"))
}

// DiscoverPort finds the specified port for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_PORT.
func DiscoverPort(service, name string) (string, error) {
	template := fmt.Sprintf(templateVar, service, name)

	return getVar(fmt.Sprintf(template, "PORT"))
}
