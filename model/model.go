package model

type Con struct {
	*LogConfig   `mapstructure:"log"`
	*MysqlConfig `mapstructure:"mysql"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	Level      string `mapstruture:"level"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	Maxbackups int    `mapstructure:"max_backups"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Port     string `mapstructure:"port"`
	Password string `mapstructure:"password"`
	MaxOpen  int    `mapstructure:"maxopen"`
	MaxFree  int    `mapstructure:"maxfree"`
}
