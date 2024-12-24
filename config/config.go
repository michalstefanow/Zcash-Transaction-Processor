package config

type Config struct {
	RPCURL      string
	RPCUser     string
	RPCPassword string
}

func LoadConfig() *Config {
	return &Config{
		RPCURL:      "http://127.0.0.1:8232",
		RPCUser:     "rpcuser",
		RPCPassword: "rpcpassword",
	}
}
