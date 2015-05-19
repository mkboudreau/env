package env

import (
	"fmt"
	"os"
)

type Environment struct {
	Readers []EnvReader
}

var (
	Docker          EnvReader = &DockerProvider{}
	CloudFoundry              = &CloudFoundryProvider{}
	Bluemix                   = &CloudFoundryProvider{}
	OpenShift                 = EnvReaderFunc(unimplemented)
	Heroku                    = EnvReaderFunc(unimplemented)
	AWS                       = EnvReaderFunc(unimplemented)
	Azure                     = EnvReaderFunc(unimplemented)
	OperatingSystem           = EnvReaderFunc(unimplemented)
)

func unimplemented(key string) string {
	return ""
}

func NewEnvironment() *Environment {
	return NewEnvironmentWithEnvReaders(Docker, CloudFoundry, OpenShift, Heroku, AWS, Azure)
}

func NewEnvironmentWithEnvReaders(readers ...EnvReader) *Environment {
	return &Environment{Readers: readers}
}

func (e *Environment) String() string {
	return fmt.Sprintf("EnvReader List: %d", e.Readers)
}

func (e *Environment) Getenv(key string) string {
	for _, reader := range e.Readers {
		if val := reader.Getenv(key); val != "" {
			return val
		}
	}

	return ""
}

func (e *Environment) SafeCopyToEnvironment(keys ...string) {
	envToSet := make(map[string]string)
	for _, key := range keys {
		if os.Getenv(key) == "" {
			if env := e.Getenv(key); env != "" {
				envToSet[key] = env
			}
		}
	}

	for key, val := range envToSet {
		os.Setenv(key, val)
	}
}

func (e *Environment) CopyAndOverrideEnvironment(keys ...string) {
	envToSet := make(map[string]string)
	for _, key := range keys {
		if env := e.Getenv(key); env != "" {
			envToSet[key] = env
		}
	}

	for key, val := range envToSet {
		if err := os.Setenv(key, val); err != nil {
			fmt.Println("Error overriding key:", key, "to val", val)
		}
	}
}
