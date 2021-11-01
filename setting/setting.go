package setting

import (
	"bluebull/model"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(model.Con)

// Init 日志初始化
func Init() (err error) {
	//配置文件的位置
	viper.SetConfigFile("./config.yml")
	//读取配置文件
	if err = viper.ReadInConfig(); err != nil {
		return model.ErrorConfRead
	}
	//把viper反序列到Conf
	if err = viper.Unmarshal(Conf); err != nil {
		return model.ErrorUnmarshalConf
	}
	//监视文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.Unmarshal(Conf)
		zap.L().Info("配置文件发生了改变")
	})
	//返回错误
	return
}
