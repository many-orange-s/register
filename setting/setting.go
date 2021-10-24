package setting

import (
	"bluebull/model"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Conf = new(model.Con)

func Init() (err error) {
	//配置文件的位置
	viper.SetConfigFile("./config.yml")
	//读取配置文件
	if err = viper.ReadInConfig(); err != nil {

		return
	}
	//把viper反序列到Conf
	if err = viper.Unmarshal(Conf); err != nil {

	}
	//监视文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {

	})
	//返回错误
	return
}
