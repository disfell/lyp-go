package config

import (
	"github.com/spf13/viper"
	"lyp-go/logger"
	"strings"
)

func init() {
	// 设置配置文件路径和类型
	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("读取配置文件失败: %s", err)
	}
	// 自动从环境变量获取配置
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func Take() {

}
