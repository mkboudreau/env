package env

import (
	"fmt"
)

type Environment struct {
	Readers []EnvReader
}

var (
	Docker EnvReader = &DockerProvider{}

	OperationSystem EnvReader = EnvReaderFunc(unimplemented)
	Bluemix         EnvReader = EnvReaderFunc(unimplemented)
	CloudFoundry    EnvReader = &CloudFoundryProvider{}
	OpenShift       EnvReader = EnvReaderFunc(unimplemented)
	Heroku          EnvReader = EnvReaderFunc(unimplemented)
	AWS             EnvReader = EnvReaderFunc(unimplemented)
	Azure           EnvReader = EnvReaderFunc(unimplemented)
)

func unimplemented(key string) string {
	return "unimplemented"
}

var DEFAULT_ENVREADER_ORDER = []EnvReader{OperationSystem, Docker, CloudFoundry, OpenShift, Heroku, AWS, Azure}

func NewEnvironment() *Environment {
	return NewEnvironmentWithEnvReaders(DEFAULT_ENVREADER_ORDER)
}

func NewEnvironmentWithEnvReaders(readers []EnvReader) *Environment {
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
