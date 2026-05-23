package config

import "github.com/halimdotnet/lulusiango/pkg/config"

const (
	DefaultConfigPath      = "./configs"
	DefaultConfigName      = "config"
	DefaultConfigExtension = "json"
)

func BindJSONConfigKey[T any](key string) (T, error) {
	return BindConfig[T](key, DefaultConfigPath, DefaultConfigName, DefaultConfigExtension)
}

func BindYAMLConfigKey[T any](key string) (T, error) {
	return BindConfig[T](key, DefaultConfigPath, DefaultConfigName, "yaml")
}

func BindConfig[T any](key, path, name, extension string) (T, error) {
	return config.Bind[T](key, path, name, extension)
}
