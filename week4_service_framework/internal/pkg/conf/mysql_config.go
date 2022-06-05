package conf

type MySQLConfig struct {
	Host      string `toml:"host"`
	Port      uint16 `toml:"port"`
	Consul    string `toml:"consul"`
	Username  string `toml:"username"`
	Password  string `toml:"password"`
	DB        string `toml:"db"`
	LogDetail bool   `toml:"log_detail"`
}
