package logger

type LoggerOption struct {
	Level         string `mapstructure:"level"`
	IsDevelopment bool   `mapstructure:"is_development"`
	EnableCaller  bool   `mapstructure:"enable_caller"`
	EnableTrace   bool   `mapstructure:"enable_trace"`
}
