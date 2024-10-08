package config

import (
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	v   *viper.Viper
	err error
)
var (
	Port string
)

// InitConfig 初始化viper配置文件
func InitConfig() {
	v = viper.New() // 构建 Viper 实例

	v.SetConfigType("yaml")   // 设置配置文件类型
	v.SetConfigName("config") // 配置文件名称(无扩展名)
	//v.SetConfigFile("./config.yaml")
	fmt.Println("-----------------")
	// 设置配置文件所在路径
	//v.AddConfigPath("./..") //./..
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
	//fmt.Printf("Using config file: %s\n", v.ConfigFileUsed())
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
	//fmt.Println("打印读取值", Port)
	//v.SetDefault(Port, ":3000")
	//fmt.Println(Port, "设置默认值")
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
func viperLoadConf() {
	Port = v.GetString("server.port")
}
