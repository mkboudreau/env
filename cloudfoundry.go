package env

import (
	"fmt"
	//"io"
	//"log"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

// CloudFoundryProvider is an implementation of a EnvReader
type CloudFoundryProvider struct{}

type VCAP struct {
	Services map[string]interface{}
}

func NewVCAPFromJSON(jsonString string) *VCAP {
	var services map[string]interface{}
	json.Unmarshal([]byte(jsonString), &services)
	return &VCAP{Services: services}
}

// Getenv looks for docker specific environment variables and returns
// either the value or empty string, if nothing is found.
func (er *CloudFoundryProvider) Getenv(key string) string {
	cfEnv := extractCloudFoundryEnvironment()
	return cfEnv[key]
}

func getValue(val interface{}) string {
	if val == nil {
		return ""
	}
	switch val.(type) {
	case string:
		return val.(string)
	case int:
		return strconv.Itoa(val.(int))
	default:
		return fmt.Sprintf("%v", val)
	}
}
func buildKey(prefix string, name string) string {
	return fmt.Sprintf("%v_%v", prefix, name)
}

func extractServices(service map[string]interface{}) map[string]string {
	serviceMap := make(map[string]string)

	prefix, ok := service["name"].(string)
	if !ok {
		fmt.Printf("Could not extract service name")
		return serviceMap
	}
	prefix = strings.ToUpper(prefix)

	serviceMap[buildKey(prefix, "LABEL")] = getValue(service["label"])
	serviceMap[buildKey(prefix, "PLAN")] = getValue(service["plan"])

	if credentialMap, ok := service["credentials"].(map[string]interface{}); ok {
		if credentialMap["username"] != nil {
			serviceMap[buildKey(prefix, "USERNAME")] = getValue(credentialMap["username"])
		}
		if credentialMap["password"] != nil {
			serviceMap[buildKey(prefix, "PASSWORD")] = getValue(credentialMap["password"])
		}
		if credentialMap["host"] != nil {
			serviceMap[buildKey(prefix, "HOST")] = getValue(credentialMap["host"])
			serviceMap[buildKey(prefix, "ADDR")] = getValue(credentialMap["host"])
		}
		if credentialMap["port"] != nil {
			serviceMap[buildKey(prefix, "PORT")] = getValue(credentialMap["port"])
		}
		if credentialMap["db"] != nil {
			serviceMap[buildKey(prefix, "NAME")] = getValue(credentialMap["db"])
		}
		if credentialMap["url"] != nil {
			serviceMap[buildKey(prefix, "URL")] = getValue(credentialMap["url"])
		}
		if credentialMap["jdbcurl"] != nil {
			serviceMap[buildKey(prefix, "JDBCURL")] = getValue(credentialMap["jdbcurl"])
			if credentialMap["url"] == nil {
				serviceMap[buildKey(prefix, "URL")] = getValue(credentialMap["jdbcurl"])
			}
		}
		if credentialMap["uri"] != nil {
			serviceMap[buildKey(prefix, "URI")] = getValue(credentialMap["uri"])
			if credentialMap["url"] == nil {
				serviceMap[buildKey(prefix, "URL")] = getValue(credentialMap["uri"])
			}
		}
	}

	return serviceMap
}

func mergeMaps(maps ...map[string]string) map[string]string {
	finalMap := make(map[string]string)
	for _, m := range maps {
		for k, v := range m {
			finalMap[k] = v
		}
	}
	return finalMap
}

func extractCloudFoundryEnvironment() map[string]string {
	normalizedEnvironment := make(map[string]string)

	vcapString := os.Getenv("VCAP_SERVICES")
	if vcapString == "" {
		return normalizedEnvironment
	}

	vcap := NewVCAPFromJSON(vcapString)

	for _, value := range vcap.Services {
		switch value.(type) {
		case map[string]interface{}:
			normalizedEnvironment = mergeMaps(normalizedEnvironment, extractServices(value.(map[string]interface{})))
		case []interface{}:
			valueItem := value.([]interface{})
			for _, slicedValueItem := range valueItem {
				if valueItemMap, ok := slicedValueItem.(map[string]interface{}); ok {
					normalizedEnvironment = mergeMaps(normalizedEnvironment, extractServices(valueItemMap))
				}
			}
		default:
			fmt.Println("--- unknown --- ")
		}
	}

	return normalizedEnvironment
}
