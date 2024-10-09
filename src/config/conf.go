package config

import "time"

// Config 是整个配置的根结构体
type Config struct {
	Service  ServiceConf  `mapstructure:"service" json:"service" yaml:"service"`
	Mysql    MysqlConf    `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis    RedisConf    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Rabbitmq RabbitmqConf `mapstructure:"rabbitmq" json:"rabbitmq" yaml:"rabbitmq"`
	Casbin   CasbinConf   `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	Etcd     EtcdConf     `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
	Click    ClickConf    `mapstructure:"clickhouse" json:"clickhouse" yaml:"clickhouse"`
	PostGre  PostGreConf  `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
	Mongo    MongoConf    `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
	Captcha  Captcha      `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Log      LogConf      `mapstructure:"log" json:"log" yaml:"log"`
}

// ServiceConf 定义服务相关的配置
type ServiceConf struct {
	Port      string `json:"port" yaml:"port" mapstructure:"port"`                   // 服务端口号
	SecretKey string `json:"secret_key" yaml:"secret_key" mapstructure:"secret_key"` // 服务秘钥
	ApiKey    string `json:"api_key" yaml:"api_key" mapstructure:"api_key"`          //api密钥验证
	CSRFKey   string `json:"csrf_key" yaml:"csrf_key" mapstructure:"csrf_key"`       //csrf请求密钥
}

// MysqlConf 定义 MySQL 数据库配置
type MysqlConf struct {
	UserName string        `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord string        `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	Host     string        `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string        `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	Database string        `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	Charset  string        `json:"charset,omitempty" yaml:"charset" mapstructure:"charset"`
	Timeout  int64         `json:"timeout,omitempty" yaml:"timeout" mapstructure:"timeout"`
	Pool     MysqlPoolConf `json:"pool" yaml:"pool" mapstructure:"pool"`
}
type MysqlPoolConf struct {
	MaxOpenConns    int `json:"max_open_conns" yaml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns    int `json:"max_idle_conns" yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	ConnMaxLifetime int `json:"conn_max_lifetime" yaml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
}

// RedisConf 定义 Redis 数据库配置
type RedisConf struct {
	UserName   string `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord   string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	Host       string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port       string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	Db         int    `json:"db,omitempty" yaml:"db" mapstructure:"db"`
	PoolSize   int    `json:"poolsize,omitempty" yaml:"poolsize" mapstructure:"poolsize"`
	MaxRetries int    `json:"max_retries,omitempty" yaml:"max_retries" mapstructure:"max_retries"`
}

// RabbitmqConf 定义 RabbitMQ 配置
type RabbitmqConf struct {
	URL      string           `json:"url,omitempty" yaml:"url" mapstructure:"url"`
	UserName string           `json:"username" yaml:"username" mapstructure:"username"`
	PassWord string           `json:"password" yaml:"password" mapstructure:"password"`
	Host     string           `json:"host" yaml:"host" mapstructure:"host"`
	Port     string           `json:"port" yaml:"port" mapstructure:"port"`
	Pool     RabbitmqPoolConf `json:"pool" yaml:"pool" mapstructure:"pool"`
}

type RabbitmqPoolConf struct {
	MaxChannels     int `json:"max_channels" yaml:"max_channels" mapstructure:"max_channels"`
	MaxIdleChannels int `json:"max_idle_channels" yaml:"max_idle_channels" mapstructure:"max_idle_channels"`
}

// CasbinConf 定义 Casbin 权限管理配置
type CasbinConf struct {
	Type     string `json:"type" yaml:"type" mapstructure:"type"`
	UserName string `json:"username,omitempty" yaml:"username" mapstructure:"username"`
	PassWord string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	Host     string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	Database string `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	Exist    bool   `json:"exist,omitempty" yaml:"exist" mapstructure:"exist"`
}

// EtcdConf 定义 Etcd 配置
type EtcdConf struct {
	Host string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
}

// ClickConf 定义 ClickHouse 配置
type ClickConf struct {
	Host     string `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	Name     string `json:"name,omitempty" yaml:"name" mapstructure:"name"`
	Password string `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	Database string `json:"database,omitempty" yaml:"database" mapstructure:"database"`
}

// LogConf 定义日志配置
type LogConf struct {
	LogPath  string `json:"log_path,omitempty" yaml:"log_path" mapstructure:"log_path"`
	LinkName string `json:"link_name,omitempty" yaml:"link_name" mapstructure:"link_name"`
	LogLevel string `json:"log_level" yaml:"log_level" mapstructure:"log_level"`
}

// Captcha 定义验证码相关配置
type Captcha struct {
	Prefix  string        `json:"prefix,omitempty" yaml:"prefix" mapstructure:"prefix"`
	Limit   int           `json:"limit" yaml:"limit" mapstructure:"limit"`
	Expired time.Duration `json:"expired,omitempty" yaml:"expired" mapstructure:"expired"`
}

// PostGreConf 定义 PostgreSQL 数据库配置
type PostGreConf struct {
	Host     string          `json:"host,omitempty" yaml:"host" mapstructure:"host"`
	Port     string          `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	Database string          `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	User     string          `json:"user,omitempty" yaml:"user" mapstructure:"user"`
	Password string          `json:"password,omitempty" yaml:"password" mapstructure:"password"`
	Pool     PostGrePoolConf `json:"pool" yaml:"pool" mapstructure:"pool"`
}

type PostGrePoolConf struct {
	MaxOpenConns    int `json:"max_open_conns" yaml:"max_open_conns" mapstructure:"max_open_conns"`
	MaxIdleConns    int `json:"max_idle_conns" yaml:"max_idle_conns" mapstructure:"max_idle_conns"`
	ConnMaxLifetime int `json:"conn_max_lifetime" yaml:"conn_max_lifetime" mapstructure:"conn_max_lifetime"`
}

// MongoConf 定义 MongoDB 配置
type MongoConf struct {
	Ip       string        `json:"ip,omitempty" yaml:"ip" mapstructure:"ip"`
	Port     string        `json:"port,omitempty" yaml:"port" mapstructure:"port"`
	Database string        `json:"database,omitempty" yaml:"database" mapstructure:"database"`
	Pool     MongoPoolConf `json:"pool" yaml:"pool" mapstructure:"pool"`
}

type MongoPoolConf struct {
	MaxPoolSize int `json:"max_pool_size" yaml:"max_pool_size" mapstructure:"max_pool_size"`
	MinPoolSize int `json:"min_pool_size" yaml:"min_pool_size" mapstructure:"min_pool_size"`
}
