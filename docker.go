package env

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// DockerProvider is an implementation of a EnvReader
type DockerProvider struct{}

// Getenv looks for docker specific environment variables and returns
// either the value or empty string, if nothing is found.
func (er *DockerProvider) Getenv(key string) string {
	dockerEnv := extractDockerEnvironment()
	return dockerEnv[key]
}

func extractDockerEnvironment() map[string]string {
	normalizedEnvironment := make(map[string]string)

	for _, entry := range os.Environ() {
		splitEntry := strings.SplitN(entry, "=", 2)
		key := splitEntry[0]
		val := splitEntry[1]

		switch {
		case strings.HasSuffix(key, "_PORT") && !strings.Contains(key, "_PORT_"):
			url := fmt.Sprintf("%v_URL", strings.TrimSuffix(key, "_PORT"))
			normalizedEnvironment[url] = val
		case strings.Contains(key, "_PORT_"):
			vars := strings.Split(key, "_PORT_")
			switch {
			case strings.Contains(key, "_PROTO"):
				proto := fmt.Sprintf("%v_PROTO", vars[0])
				normalizedEnvironment[proto] = val
			case strings.Contains(key, "_ADDR"):
				addr := fmt.Sprintf("%v_ADDR", vars[0])
				host := fmt.Sprintf("%v_HOST", vars[0])
				normalizedEnvironment[addr] = val
				normalizedEnvironment[host] = val
			case strings.Contains(key, "_PORT"):
				port := fmt.Sprintf("%v_PORT", vars[0])
				normalizedEnvironment[port] = val
			}
		case strings.Contains(key, "_ENV_"):
			vars := strings.Split(key, "_ENV_")
			if len(vars) != 2 {
				log.Printf("length of %v is not equal to 2 after split on _ENV_", key)
			}
			envKey := fmt.Sprintf("%v_%v", vars[0], vars[1])
			normalizedEnvironment[envKey] = val
		}
	}

	return normalizedEnvironment
}

/*
	{map[string]string{
		"DB_NAME":                     "/somename/db",
		"DB_PORT":                     "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":             "TEST123",
	},
		[]string{"", "DB_PORT", "DB_HOST", "DB_ADDR", "DB_URL", "DB_ANYTHING"},
		[]string{"", "5555", "24.50.222.1", "24.50.222.1", "tcp://24.50.222.1:5555", "TEST123"},
	},
*/
