package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// 用来读取配置文件
func InitConfig() {
	workDir, _ := os.Getwd()                 // 工作目录的路径
	viper.SetConfigName("config")            // 配置文件的文件名
	viper.SetConfigType("yml")               // 配置文件的后缀
	viper.AddConfigPath(workDir + "/config") // 获取到配置文件的路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取数据库配置错误")
		panic(err)
	}
}
