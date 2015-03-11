// Package cloud-env provides an API consistent way to retrieve environment vars from popular cloud hosting providers
// such as Heroku, Open Shift, Bluemix & Cloud Foundry, and Docker.
// It creates a default search-order precedent that is applicable in most cases. However, if needed, it can be overridden.
package env

// Interface with one function that is identical to the standard library's os.Getenv
type EnvReader interface {
	Getenv(key string) string
}

type EnvReaderFunc func(key string) string

func (erFunc EnvReaderFunc) Getenv(key string) string {
	return erFunc(key)
}
