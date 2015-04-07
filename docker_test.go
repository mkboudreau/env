package env

import (
	"os"
	"testing"
)

var dockerTestCases = []struct {
	setup    map[string]string
	test     []string
	expected []string
}{
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{""},
		[]string{""},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{"DB_PORT"},
		[]string{"5555"},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{"DB_HOST"},
		[]string{"24.50.222.1"},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{"DB_ADDR"},
		[]string{"24.50.222.1"},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{"DB_URL"},
		[]string{"tcp://24.50.222.1:5555"},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{"DB_ANYTHING"},
		[]string{"TEST123"},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{},
		[]string{},
		//[]string{"", "DB_PORT", "DB_HOST", "DB_ADDR", "DB_URL", "DB_ANYTHING"},
		//[]string{"", "5555", "24.50.222.1", "24.50.222.1", "tcp://24.50.222.1:5555", "TEST123"},
	},
	{map[string]string{
		"DB_NAME":                "/somename/db",
		"DB_PORT":                "tcp://24.50.222.1:5555",
		"DB_PORT_5555_TCP_PROTO": "tcp",
		"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
		"DB_PORT_5555_TCP_PORT":  "5555",
		"DB_ENV_ANYTHING":        "TEST123",
	},
		[]string{},
		[]string{},
	},
}

func setupEnvironmentWithMap(envMap map[string]string) {
	for key, val := range envMap {
		os.Setenv(key, val)
	}
}

func clearEnvironmentWithMap(envMap map[string]string) {
	for key := range envMap {
		os.Unsetenv(key)
	}
}

func TestDockerProviderTaskRunner(t *testing.T) {
	provider := &DockerProvider{}

	for _, testCase := range dockerTestCases {
		setupEnvironmentWithMap(testCase.setup)

		if IsNotEqual(provider, testCase.test, testCase.expected) {
			t.Errorf("Failure: actual %v does not equal expected %v", testCase.test, testCase.expected)
		} else {
			t.Logf("Success: actual %v does equal expected %v", testCase.test, testCase.expected)
		}

		clearEnvironmentWithMap(testCase.setup)
	}
}

func IsNotEqual(reader EnvReader, keys []string, test []string) bool {
	return !IsEqual(reader, keys, test)
}
func IsEqual(reader EnvReader, keys []string, test []string) bool {
	if len(keys) != len(test) {
		return false
	}

	for i := range keys {
		if reader.Getenv(keys[i]) != test[i] {
			return false
		}
	}

	return true
}
