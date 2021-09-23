package configuration

type ServerConfiguration struct {
	Port             string `toml:"port"`
	APIVersion       string `toml:"api_version"`
	Debug            bool   `env:"SERVER_DEBUG"`
	UserPasswordSalt string `env:"SERVER_USER_PASSWORD_SALT"`
}

type DatabaseConfiguration struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	DbName   string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SslMode  string `env:"DB_SSLMODE"`
}

type JwtConfiguration struct {
	SigningKey      string `env:"JWT_SIGNING_KEY"`
	AccessTokenTTL  string `env:"JWT_ACCESS_TOKEN_TTL"`
	RefreshTokenTTL string `env:"JWT_REFRESH_TOKEN_TTL"`
}

type Configuration struct {
	ServerConfiguration ServerConfiguration `toml:"server"`
	DbConfiguration     DatabaseConfiguration
	AuthConfiguration   JwtConfiguration
}
