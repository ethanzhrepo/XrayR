package limiter

type GlobalDeviceLimitConfig struct {
	Enable        bool   `mapstructure:"Enable"`
	RedisNetwork  string `mapstructure:"RedisNetwork"` // tcp or unix
	RedisAddr     string `mapstructure:"RedisAddr"`    // host:port, or /path/to/unix.sock
	RedisUsername string `mapstructure:"RedisUsername"`
	RedisPassword string `mapstructure:"RedisPassword"`
	RedisDB       int    `mapstructure:"RedisDB"`

	RedisTLSEnabled    bool   `mapstructure:"RedisTLSEnabled"`    // enable tls, must host:port
	RedisCACert        string `mapstructure:"RedisCACert"`        // ca certification
	RedisClientCert    string `mapstructure:"RedisClientCert"`    // client certification
	RedisClientCertKey string `mapstructure:"RedisClientCertKey"` // client certification

	Timeout int `mapstructure:"Timeout"`
	Expiry  int `mapstructure:"Expiry"` // second
}
