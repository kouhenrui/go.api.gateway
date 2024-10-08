package database

import (
	_ "database/sql"
	"fmt"
	"go.api.gateway/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var (
	MysqlClient *gorm.DB
	err         error
)

//type MysqlT struct {
//	client *gorm.DB
//}

func NewMysqlClient(mysqlConfig config.MysqlConf) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", mysqlConfig.UserName, mysqlConfig.PassWord, mysqlConfig.Host, mysqlConfig.Database, mysqlConfig.Charset)
	//dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", MysqlConfig.UserName, MysqlConfig.PassWord, MysqlConfig.HOST, MysqlConfig.DATABASE, MysqlConfig.CHARSET) //&timeout=%s , MysqlConfig.TimeOut
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	MysqlClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, //跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //false 复数形式
			//TablePrefix:   "",    //表名前缀 User的表名应该是t_users
		},
		DisableForeignKeyConstraintWhenMigrating: true, //设置成为逻辑外键(在物理数据库上没有外键，仅体现在代码上)

	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := MysqlClient.DB()
	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100)                 //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)                  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	sqlDB.SetConnMaxIdleTime(30 * time.Minute) //设置30秒重连

	// 设置重试逻辑
	//retryCount := 5
	//MysqlClient.WithContext(context.Background()).Retry(retryCount, time.Second, func() error {
	//	// 尝试连接数据库
	//	dbSQL, erro := MysqlClient.DB()
	//	if erro != nil {
	//		return erro
	//	}
	//
	//	return dbSQL.Ping()
	//})
	log.Printf("mysql初始化连接成功")
	initTable()
	log.Println("数据库表结构初始化成功")
}

// var (
//
//	acc = pojo.Account{}
//	apf = pojo.AccountProfile{}
//	mer = pojo.Merchant{}
//	cat = pojo.Category{}
//	rol = pojo.Role{}
//	pro = pojo.Product{}
//	dep = pojo.Department{}
//
// )

// 迁移表
func initTable() {
	err = MysqlClient.AutoMigrate()
	if err != nil {
		log.Println("迁移表失败")
	} //acc, mer, cat, rol, pro, apf, dep
}

//func NewSqlClient() *MysqlT {
//	return &MysqlT{client: MysqlClient}
//}
