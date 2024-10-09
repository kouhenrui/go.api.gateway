package database

import (
	"fmt"
	"go.api.gateway/src/config"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"time"
)

var (
	CKClient *gorm.DB
)

func InitClickhouse(conf config.ClickConf) {
	dsn := fmt.Sprintf("clickhouse://gorm:gorm@%s:%s/%sdial_timeout=10s&read_timeout=20s", conf.Host, conf.Port, conf.Database)
	//dsn := "clickhouse://gorm:gorm@localhost:9942/gorm?dial_timeout=10s&read_timeout=20s"
	CKClient, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})

	sqldb, _ := CKClient.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqldb.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqldb.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqldb.SetConnMaxLifetime(time.Hour)
}
