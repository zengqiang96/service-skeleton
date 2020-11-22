package config

// Config 全局的配置
type Config struct {
	Network string `json:"network"`
	Address string `json:"address"`
}

// Load 加载配置文件
func Load(configFile string) (*Config, error) {
	return nil, nil
}
