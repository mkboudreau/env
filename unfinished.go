package env

import (
	"os"
)

type OperatingSystemProvider struct{}

func (er *OperatingSystemProvider) Getenv(key string) string {
	return os.Getenv(key)
}

type OpenShiftProvider struct{}

func (er *OpenShiftProvider) Getenv(key string) string {
	return ""
}

type HerokuProvider struct{}

func (er *HerokuProvider) Getenv(key string) string {
	return ""
}

type AWSProvider struct{}

func (er *AWSProvider) Getenv(key string) string {
	return ""
}

type AzureProvider struct{}

func (er *AzureProvider) Getenv(key string) string {
	return ""
}
