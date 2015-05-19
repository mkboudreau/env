package env

import (
	"fmt"
	"os"
	"testing"
)

func TestAllProviders(t *testing.T) {
	setupEnvironmentTest(integrationTestSetup)
	environment := NewEnvironment()
	RunTestCasesWithEnvReader(t, integrationTestCases, environment)
}

func TestSafeCopyToOS(t *testing.T) {
	environment := NewEnvironment()
	postEnvSetup := func() {
		environment.SafeCopyToEnvironment("DB_PORT", "DB_HOST", "DB_ADDR", "DB_URL", "DB_USERNAME", "DB1_PORT", "DB1_HOST", "DB1_ADDR", "DB1_URL", "DB1_USERNAME")
	}
	RunTestCasesWithEnvReaderWithPostEnvSetup(t, integrationTestCases, EnvReaderFunc(os.Getenv), postEnvSetup)
}

func TestOverrideOS(t *testing.T) {
	environment := NewEnvironment()
	postEnvSetup := func() {
		environment.CopyAndOverrideEnvironment("DB_PORT", "DB_HOST", "DB_ADDR", "DB_URL", "DB_USERNAME", "DB1_PORT", "DB1_HOST", "DB1_ADDR", "DB1_URL", "DB1_USERNAME")
	}
	RunTestCasesWithEnvReaderWithPostEnvSetup(t, integrationTestCases, EnvReaderFunc(os.Getenv), postEnvSetup)
}

func TestSafeCopyConflictToUs(t *testing.T) {
	environment := NewEnvironment()
	postEnvSetup := func() {
		environment.SafeCopyToEnvironment("DB_PORT", "DB_URL")
	}
	RunTestCasesWithEnvReaderWithPostEnvSetup(t, dbportWithSafeCopy, EnvReaderFunc(os.Getenv), postEnvSetup)
}

func TestOverrideConflictOS(t *testing.T) {
	environment := NewEnvironment()
	postEnvSetup := func() {
		environment.CopyAndOverrideEnvironment("DB_PORT", "DB_URL")
	}
	RunTestCasesWithEnvReaderWithPostEnvSetup(t, dbportWithOverride, EnvReaderFunc(os.Getenv), postEnvSetup)
}

func setupEnvironmentTest(envvars map[string]string) {
	for key, value := range envvars {
		os.Setenv(key, value)
	}
}
func debugEnvironment() {
	for _, value := range os.Environ() {
		fmt.Println(">>> DEBUG >>> Env Entry:", value)
	}
}

var integrationTestSetup map[string]string = map[string]string{
	"DB_NAME":                "/somename/db",
	"DB_PORT":                "tcp://24.50.222.1:5555",
	"DB_PORT_5555_TCP_PROTO": "tcp",
	"DB_PORT_5555_TCP_ADDR":  "24.50.222.1",
	"DB_PORT_5555_TCP_PORT":  "5555",
	"DB_ENV_ANYTHING":        "TEST123",
	"VCAP_SERVICES": `
		{
		  "sqldb": {
		    "name": "localdb",
		    "label": "sqldb",
		    "plan": "sqldb_small",
		    "credentials": {
		      "port": 5000,
		      "db": "I_04505",
		      "username": "baaavkd",
		      "host": "3.24.23.46",
		      "hostname": "3.24.23.46",
		      "jdbcurl": "jdbc:db2://3.24.23.46:5000/I_04505",
		      "uri": "db2://bgjdkw:uzhudznsvlc@3.24.33.46:5000/I_04505",
		      "password": "uzhuvjznsvlc"
		    }
		  },
		  "cloudantNoSQLDB": [
		      {
			         "name": "db1",
			         "label": "cloudantNoSQLDB",
			         "plan": "Shared",
			         "credentials": {
				     "username": "bgjeavkw",
				     "password": "uzhuvjznsvlc",
			         "host": "9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com",
			         "port": 443,
			         "url": "https://user:pass@9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"
		         }
		      }
		   ]
		}`,
}

var dbportWithOverride = []envReaderTestCases{
	{integrationTestSetup,
		[]string{"DB_PORT"},
		[]string{"5555"},
	},
	{integrationTestSetup,
		[]string{"DB_URL"},
		[]string{"tcp://24.50.222.1:5555"},
	},
}
var dbportWithSafeCopy = []envReaderTestCases{
	{integrationTestSetup,
		[]string{"DB_PORT"},
		[]string{"tcp://24.50.222.1:5555"},
	},
	{integrationTestSetup,
		[]string{"DB_URL"},
		[]string{"tcp://24.50.222.1:5555"},
	},
}
var integrationTestCases = []envReaderTestCases{
	{integrationTestSetup,
		[]string{},
		[]string{},
	},
	{integrationTestSetup,
		[]string{"DB_HOST"},
		[]string{"24.50.222.1"},
	},
	{integrationTestSetup,
		[]string{"DB_ADDR"},
		[]string{"24.50.222.1"},
	},
	{integrationTestSetup,
		[]string{"DB_URL"},
		[]string{"tcp://24.50.222.1:5555"},
	},
	{integrationTestSetup,
		[]string{"DB_USERNAME"},
		[]string{""},
	},
	{integrationTestSetup,
		[]string{"DB1_PORT"},
		[]string{"443"},
	},
	{integrationTestSetup,
		[]string{"DB1_HOST"},
		[]string{"9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"},
	},
	{integrationTestSetup,
		[]string{"DB1_ADDR"},
		[]string{"9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"},
	},
	{integrationTestSetup,
		[]string{"DB1_URL"},
		[]string{"https://user:pass@9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"},
	},
	{integrationTestSetup,
		[]string{"DB1_USERNAME"},
		[]string{"bgjeavkw"},
	},
}
