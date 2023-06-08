package config

import (
	"github.com/spf13/viper"
	"os"
)

func LoadConfig() {

	// 初始化 viper
	// 检查配置文件是否存在
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		// 配置文件不存在，绑定环境变量
		err := viper.BindEnv("MYSQL_HOST")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_PORT")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_USER")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_PASSWORD")
		if err != nil {
			return
		}
		err = viper.BindEnv("MYSQL_DB")
		if err != nil {
			return
		}
	} else {
		// 配置文件存在，读取配置文件
		viper.SetConfigFile(".env") // 设置配置文件名，可以是 .yaml、.toml、.json 等格式

		viper.SetDefault("PROJECT_NAME", "jsy_customer") // 设置默认值
		viper.SetDefault("PROJECT_DESCRIPTION", "jsy_customer")
		viper.SetDefault("PROJECT_VERSION", "v0.0.1")
		viper.SetDefault("DEBUG", false)
		viper.SetDefault("MYSQL_HOST", "localhost")
		viper.SetDefault("MYSQL_PORT", "3306")
		viper.SetDefault("MYSQL_USER", "user")
		viper.SetDefault("MYSQL_PASSWORD", "password")
		viper.SetDefault("MYSQL_DB", "mysql")

		if err := viper.ReadInConfig(); err != nil {
			// 读取配置文件失败
			panic("Error reading config file: " + err.Error())
		}
	}

}
