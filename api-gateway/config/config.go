package config

import (
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()                 // 工作目录的路径
	viper.SetConfigName("config")            // 配置文件的文件名
	viper.SetConfigType("yml")               // 配置文件后缀
	viper.AddConfigPath(workDir + "/config") // 获取到配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {

	}
}
