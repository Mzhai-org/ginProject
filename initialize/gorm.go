package initialize

import (
	"fmt"
	"ginProject/global"
	"ginProject/initialize/internal"
	"ginProject/model"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

func Gorm() *gorm.DB {
	return GormMysql()
}

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	//global.GVA_LOG.Info("数据库连接信息", zap.String("dsn", dsn))\
	fmt.Printf("数据库连接信息", dsn)
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig(m.LogMode)); err != nil {
		global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		global.GVA_LOG.Info("MySQL启动成功")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func gormConfig(mod bool) *gorm.Config {
	var config = &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Mysql.LogZap {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	case "zap", "Zap":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		if mod {
			config.Logger = internal.Default.LogMode(logger.Info)
			break
		}
		config.Logger = internal.Default.LogMode(logger.Silent)
	}
	return config
}

func MysqlTables(db *gorm.DB) {
	global.GVA_LOG.Info("根据model正向生成数据库表")
	err := db.AutoMigrate(
		model.Customer{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Any("err", err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
