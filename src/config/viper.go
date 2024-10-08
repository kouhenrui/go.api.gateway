package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

var (
	v           *viper.Viper
	err         error
	ViperConfig Config
)

// InitConfig 初始化viper配置文件
func InitConfig() {
	v = viper.New()           // 构建 Viper 实例
	v.SetConfigType("yaml")   // 设置配置文件类型
	v.SetConfigName("config") // 配置文件名称(无扩展名)
	// 添加配置文件的搜索路径
	v.AddConfigPath("../")   // 搜索父目录 (相对于 conf.go 文件的位置)
	v.AddConfigPath("./src") // 搜索 src 目录
	v.AddConfigPath(".")     // 如果程序的工作目录就是 go.apo.gateway
	// 读取配置文件
	if err = v.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			fmt.Println(err, "文件未找到", err.Error())
			// 配置文件未找到
			panic(fmt.Errorf("配置文件未找到: %w", err))
			//return nil, fmt.Errorf("配置文件未找到: %w", err)
		} else {
			fmt.Println(err, "读取配置文件错误")
			panic(fmt.Errorf("读取配置文件错误: %w", err))
			//return nil, fmt.Errorf("读取配置文件错误: %w", err)
		}
	}

	viperLoadConf()
	v.WatchConfig() //开启监听
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file updated.")
		viperLoadConf() // 加载配置的方法
	})
	fmt.Println("配置初始化成功")
}

// 加载配置信息
func viperLoadConf() {
	if err = v.Unmarshal(&ViperConfig); err != nil {
		log.Fatalf("Unable to decode into config struct: %v", err)
	}
	log.Println(ViperConfig.Service.WhiteUrl, "+++++++++++++")
	InitLogger(ViperConfig.Log) // 初始化日志记录
}
