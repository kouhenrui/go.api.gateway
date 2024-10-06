package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"strings"
)

var (
	v   *viper.Viper
	err error
)

// InitConfig 初始化viper配置文件
func InitConfig() {
	v = viper.New() // 构建 Viper 实例
	//configPath := ""
	// 设置配置文件名（不带扩展名）
	v.SetConfigName("config.dev")

	// 设置配置文件类型
	v.SetConfigType("yaml")
	v.AddConfigPath("../../")
	// 如果传入了配置路径，则设置搜索路径
	//if configPath != "" {
	//	v.AddConfigPath(configPath)
	//} else {
	//	// 默认在当前目录下找配置
	//	v.AddConfigPath(".")
	//}

	// 读取环境变量
	v.SetEnvPrefix("myapp") // 设置环境变量前缀
	v.AutomaticEnv()        // 自动从环境变量读取
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// 配置文件未找到
			panic(fmt.Errorf("配置文件未找到: %w", err))
			//return nil, fmt.Errorf("配置文件未找到: %w", err)
		} else {
			// 其他读取配置错误
			panic(fmt.Errorf("读取配置文件错误: %w", err))
			//return nil, fmt.Errorf("读取配置文件错误: %w", err)
		}
	}

	// 设置默认值（如果配置文件中没有，使用这些默认值）
	//viper.SetDefault("server.port", 8080)

	// 序列化配置到结构体
	//if err := viper.Unmarshal(&Conf); err != nil {
	//	return nil, fmt.Errorf("反序列化配置文件错误: %w", err)
	//}

	// 日志初始化
	logFile := v.GetString("log_file")
	if logFile != "" {
		initLogging(logFile)
	}

	viperLoadConf()
	v.WatchConfig() //开启监听
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file updated.")
		viperLoadConf() // 加载配置的方法
	})
	fmt.Println("配置初始化成功")
	//return Conf, nil
}

// 初始化日志记录
func initLogging(logFile string) {
	// 打开日志文件
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("无法打开日志文件: %v", err)
	}
	// 设置日志输出到文件
	log.SetOutput(file)
	log.Println("日志已启动")
}

// 加载配置信息
func viperLoadConf() {}
