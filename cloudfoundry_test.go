package env

import (
	"testing"
)

func TestCloudFoundryProvider(t *testing.T) {
	RunTestCasesWithEnvReader(t, cfTestCases, &CloudFoundryProvider{})
}

var basicCfTestSetup map[string]string = map[string]string{
	"VCAP_SERVICES": `
		{
		  "sqldb": {
		    "name": "sherlock-db",
		    "label": "sqldb",
		    "plan": "sqldb_small",
		    "credentials": {
		      "port": 5000,
		      "db": "I_04505",
		      "username": "baaavkd",
		      "host": "3.24.23.46",
		      "hostname": "3.24.23.46",
		      "jdbcurl": "jdbc:db2://3.24.23.46:50000/I_04505",
		      "uri": "db2://bgjdkw:uzhudznsvlc@3.24.33.46:5000/I_04505",
		      "password": "uzhuvjznsvlc"
		    }
		  },
		  "cloudantNoSQLDB": [
		      {
		         "name": "SensorsData",
		         "label": "cloudantNoSQLDB",
		         "plan": "Shared",
		         "credentials": {
				      "username": "bgjeavkw",
				      "password": "uzhuvj",
			          "host": "9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com",
			          "port": 443,
			          "url": "https://user:pass@9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"
		         }
		      }
		   ]
		}`,
}
var cfTestCases = []envReaderTestCases{
	{basicCfTestSetup,
		[]string{},
		[]string{},
	},
	{basicCfTestSetup,
		[]string{"SHERLOCK-DB_PORT"},
		[]string{"5000"},
	},
	{basicCfTestSetup,
		[]string{"SHERLOCK-DB_HOST"},
		[]string{"3.24.23.46"},
	},
	{basicCfTestSetup,
		[]string{"SHERLOCK-DB_ADDR"},
		[]string{"3.24.23.46"},
	},
	{basicCfTestSetup,
		[]string{"SHERLOCK-DB_URL"},
		[]string{"db2://bgjdkw:uzhudznsvlc@3.24.33.46:5000/I_04505"},
	},
	{basicCfTestSetup,
		[]string{"SHERLOCK-DB_USERNAME"},
		[]string{"baaavkd"},
	},
	{basicCfTestSetup,
		[]string{"SHERLOCK-DB_PASSWORD"},
		[]string{"uzhuvjznsvlc"},
	},

	{basicCfTestSetup,
		[]string{"SENSORSDATA_PORT"},
		[]string{"443"},
	},
	{basicCfTestSetup,
		[]string{"SENSORSDATA_HOST"},
		[]string{"9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"},
	},
	{basicCfTestSetup,
		[]string{"SENSORSDATA_ADDR"},
		[]string{"9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"},
	},
	{basicCfTestSetup,
		[]string{"SENSORSDATA_URL"},
		[]string{"https://user:pass@9ac6fd0e-c143-4c8a-8484-913f36b18fff-bluemix.cloudant.com"},
	},
	{basicCfTestSetup,
		[]string{"SENSORSDATA_USERNAME"},
		[]string{"bgjeavkw"},
	},
	{basicCfTestSetup,
		[]string{"SENSORSDATA_PASSWORD"},
		[]string{"uzhuvj"},
	},
}
