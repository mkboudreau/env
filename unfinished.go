package env

type OperationSystemProvider struct{}

func (er *OperationSystemProvider) Getenv(key string) string {
	return "unimplemented"
}

type CloudFoundryProvider struct{}

func (er *CloudFoundryProvider) Getenv(key string) string {
	return "unimplemented"
}

type OpenShiftProvider struct{}

func (er *OpenShiftProvider) Getenv(key string) string {
	return "unimplemented"
}

type HerokuProvider struct{}

func (er *HerokuProvider) Getenv(key string) string {
	return "unimplemented"
}

type AWSProvider struct{}

func (er *AWSProvider) Getenv(key string) string {
	return "unimplemented"
}

type AzureProvider struct{}

func (er *AzureProvider) Getenv(key string) string {
	return "unimplemented"
}
