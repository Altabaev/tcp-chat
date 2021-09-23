package app

// Config структура конфигурации
type Config struct {
	Protocol string `toml:"protocol"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

// NewConfig конструктор конфигурации со значениями по-умолчанию
func NewConfig() *Config {
	return &Config{
		Protocol: "tcp",
		Host:     "localhost",
		Port:     "9090",
	}
}
