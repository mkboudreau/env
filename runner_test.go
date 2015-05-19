package env

import (
	"os"
	"reflect"
	"testing"
)

type envReaderTestCases struct {
	setup    map[string]string
	test     []string
	expected []string
}

func SetupEnvironmentWithMap(envMap map[string]string) {
	for key, val := range envMap {
		os.Setenv(key, val)
	}
}

func ClearEnvironmentWithMap(envMap map[string]string) {
	for key := range envMap {
		os.Unsetenv(key)
	}
}

func RunTestCasesWithEnvReader(t *testing.T, testCases []envReaderTestCases, reader EnvReader) {
	RunTestCasesWithEnvReaderWithPostEnvSetup(t, testCases, reader, nil)
}

func RunTestCasesWithEnvReaderWithPostEnvSetup(t *testing.T, testCases []envReaderTestCases, reader EnvReader, postEnvSetup func()) {
	for _, testCase := range testCases {
		SetupEnvironmentWithMap(testCase.setup)

		if postEnvSetup != nil {
			postEnvSetup()
		}

		if ok, actual := IsEqual(reader, testCase.test, testCase.expected); ok {
			t.Logf("Success [%v]: key %v does equal expected value %v", reflect.TypeOf(reader), testCase.test, testCase.expected)
		} else {
			t.Errorf("Failure [%v]: key %v does not equal expected value %v; actual %v", reflect.TypeOf(reader), testCase.test, testCase.expected, actual)
		}

		ClearEnvironmentWithMap(testCase.setup)
	}
}

func IsEqual(reader EnvReader, keys []string, test []string) (bool, string) {
	if len(keys) != len(test) {
		return false, ""
	}

	for i := range keys {
		if reader.Getenv(keys[i]) != test[i] {
			return false, reader.Getenv(keys[i])
		}
	}

	return true, ""
}
