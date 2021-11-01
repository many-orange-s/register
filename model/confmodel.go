package model

/*这个里面只放配置文件的结构体参数*/

// Con mapstructure 是反序列话的那个库
type Con struct {
	Mode         string `mapstructure:"mode"`
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	Level      string `mapstructure:"level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	Maxbackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
	MaxOpen  int    `mapstructure:"maxopen"`
	MaxFree  int    `mapstructure:"maxfree"`
}
