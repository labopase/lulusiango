package healthz

type HealthOptions struct {
	Redis    bool `mapstructure:"redis"`
	Postgres bool `mapstructure:"postgres"`
	Kafka    bool `mapstructure:"kafka"`
	Grpc     bool `mapstructure:"grpc"`
}
