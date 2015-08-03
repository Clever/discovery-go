package discovery

import (
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

func getVar(envVar string) string {
	envVar = strings.ToUpper(envVar)
	val := os.Getenv(envVar)
	if val == "" {
		panic(kv.FormatLog("discovery-go", kv.Error, "missing env var", m{
			"var": envVar,
		}))
	}
	return val
}

// Discover finds the specified URL for a service based off of the service's name and which
// interface you are accessing. Values are found in environment variables fitting the scheme:
// SERVICE_{SERVICE NAME}_{INTERFACE NAME}_{PROTO,HOST,PORT}.
func Discover(service, name string) string {
	template := fmt.Sprintf(templateVar, service, name)

	proto := getVar(fmt.Sprintf(template, "PROTO"))
	host := getVar(fmt.Sprintf(template, "HOST"))
	port := getVar(fmt.Sprintf(template, "PORT"))

	u := url.URL{
		Scheme: proto,
		Host:   fmt.Sprintf("%s:%s", host, port),
	}
	return u.String()
}
