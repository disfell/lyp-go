package config

import (
	"github.com/spf13/viper"
	"lyp-go/logger"
)

func Init() {
	// 设置配置文件路径和类型
	viper.SetConfigName("config") // 配置文件名（不带扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath(".")      // 配置文件路径

	// 读取配置文件
	err := viper.ReadInConfig()
	if err != nil {
		logger.Errorf("读取配置文件失败: %s", err)
	}

	_ = viper.BindEnv("supabase.url", "SUPABASE_URL")
	_ = viper.BindEnv("supabase.annokey", "SUPABASE_ANNOKEY")
	_ = viper.BindEnv("supabase.key", "SUPABASE_KEY")
	_ = viper.BindEnv("steam.id", "STEAM_ID")
	_ = viper.BindEnv("steam.token", "STEAM_TOKEN")
	// 自动从环境变量获取配置
	viper.AutomaticEnv()
}
