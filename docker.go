package env

type DockerProvider struct{}

func (er *DockerProvider) Getenv(key string) string {
	return "docker unimplemented"
}
