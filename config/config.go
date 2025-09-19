package config

type NebulaConfig struct {
	Username  string       `yaml:"username" json:"username"`
	Password  string       `yaml:"password" json:"password"`
	SpaceName string       `yaml:"space-name" json:"space_name"`
	Pool      PoolConfig   `yaml:"pool" json:"pool"`
	Hosts     []NebulaHost `yaml:"hosts" json:"hosts"`
}

type NebulaHost struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

// PoolConfig 连接池配置
type PoolConfig struct {
	MaxSize  int `yaml:"max-size" json:"max_size"`
	MinSize  int `yaml:"min-size" json:"min_size"`
	Timeout  int `yaml:"timeout" json:"timeout"`
	IdleTime int `yaml:"idle_time" json:"idle_time"`
}
