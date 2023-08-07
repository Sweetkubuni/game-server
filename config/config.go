package config

type IConfig interface {
	GetPort() string
	GetRedisConfig() RedisConfig
	GetCors() CorsConfig
	GetServerConfig() ServerConfig
}

// IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,

type ServerConfig struct {
	IdleTimeoutMult int
	ReadTimeoutMult int
	WriteTimeout    int
	JwtSecret       []byte
}

type CorsConfig struct {
	TrustedOrigins []string
}

type RedisConfig struct {
	Addr     string
	Password string
	Username string
}
