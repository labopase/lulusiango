package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func Bind[T any](key, path, name, extension string) (T, error) {
	var result T

	v := viper.New()

	v.SetConfigName(name)
	v.AddConfigPath(path)
	v.SetConfigType(extension)

	if err := v.ReadInConfig(); err != nil {
		return result, fmt.Errorf("viper failed to read config file: %w", err)
	}

	if key == "" {
		if err := v.Unmarshal(&result); err != nil {
			return result, fmt.Errorf("viper failed to unmarshal config file: %w", err)
		}
	} else {
		if err := v.UnmarshalKey(key, &result); err != nil {
			return result, fmt.Errorf("viper unable to decode into struct, %w", err)
		}
	}
	return result, nil
}
