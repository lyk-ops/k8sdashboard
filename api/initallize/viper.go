package initallize

import (
	"github.com/spf13/viper"
	"kubeimook/global"
)

func Viper() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")

	// 读取配置文件
	if err := v.ReadInConfig(); err != nil {
		panic(err) // 在实际项目中，您可能需要更优雅的错误处理方式
	}

	// 将配置文件内容解析到 global.CONF
	if err := v.Unmarshal(&global.CONF); err != nil {
		panic(err) // 同样，这里需要更优雅的错误处理
	}
}
