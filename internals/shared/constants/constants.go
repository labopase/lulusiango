package constants

const (
	EnvDevelopment = "development"
	EnvTest        = "test"
	EnvProduction  = "production"
)

var ListEnv = []string{
	EnvDevelopment,
	EnvTest,
	EnvProduction,
}

const (
	LogInfo  = "info"
	LogError = "error"
	LogDebug = "debug"
	LogFatal = "fatal"
	LogPanic = "panic"
	LogWarn  = "warn"
)

const (
	RequestTenantID = "tenant_id"
	RequestUserID   = "user_id"
	RequestAPIKeyID = "api_key_id"
)
