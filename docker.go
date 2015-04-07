package env

import (
	"fmt"
	"os"
	"strings"
)

type DockerProvider struct{}

func (er *DockerProvider) Getenv(key string) string {
	return "docker unimplemented"
}

func translateEnvsToExpected(key string) ([]string, error) {
	envs := make([]string, 10)
	for _, val := range os.Environ() {
		switch {
		case strings.HasSuffix(val, "_PORT"):
			envs = append(envs, fmt.Sprintf("%vURL", strings.TrimSuffix(val, "_PORT")))
		case strings.Contains(val, "_PORT_"):
		case strings.Contains(val, "_ENV_"):
			vars := strings.Split(val, "_ENV_")
			if len(vars) != 2 {
				return nil, fmt.Errorf("length of %v is not equal to 2 after split on _ENV_", val)
			}
			envs = append(envs, fmt.Sprintf("%v_%v", vars[0], vars[1]))

		}
	}

	return envs, nil
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
