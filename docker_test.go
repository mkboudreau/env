package env

import (
	"testing"
)

func TestDockerProvider(t *testing.T) {
	RunTestCasesWithEnvReader(t, dockerTestCases, &DockerProvider{})
}

var basicDockerTestSetup map[string]string = map[string]string{
	"DB_NAME":                "/somename/db",
	"DB_PORT":                "tcp://24.50.222.1:5555",
	"DB_PORT_5555_TCP_PROTO": "tcp",
	"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
	"DB_PORT_5555_TCP_PORT":  "5555",
	"DB_ENV_ANYTHING":        "TEST123",
}
var dockerTestCases = []envReaderTestCases{
	{basicDockerTestSetup,
		[]string{""},
		[]string{""},
	},
	{basicDockerTestSetup,
		[]string{"DB_PORT"},
		[]string{"5555"},
	},
	{basicDockerTestSetup,
		[]string{"DB_HOST"},
		[]string{"24.50.222.1"},
	},
	{basicDockerTestSetup,
		[]string{"DB_ADDR"},
		[]string{"24.50.222.1"},
	},
	{basicDockerTestSetup,
		[]string{"DB_URL"},
		[]string{"tcp://24.50.222.1:5555"},
	},
	{basicDockerTestSetup,
		[]string{"DB_ANYTHING"},
		[]string{"TEST123"},
	},
	{basicDockerTestSetup,
		[]string{"", "DB_PORT", "DB_HOST", "DB_ADDR", "DB_URL", "DB_ANYTHING"},
		[]string{"", "5555", "24.50.222.1", "24.50.222.1", "tcp://24.50.222.1:5555", "TEST123"},
	},
	{basicDockerTestSetup,
		[]string{},
		[]string{},
	},
}
