package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test/week4_service_framework/internal/pkg/conf"
)

var DB *gorm.DB

func NewDB(mysqlConfig *conf.MySQLConfig) *gorm.DB {
	var (
		err         error
		engine      *gorm.DB
		_sql        *sql.DB
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		mysqlConfig.Username,
		mysqlConfig.Password,
		mysqlConfig.Host,
		mysqlConfig.Port,
		mysqlConfig.DB,
	)

	if engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(fmt.Errorf("failed to connect mysql, %w", err))
	}
	if _sql, err = engine.DB(); err != nil {
		panic(fmt.Errorf("failed to connect mysql, %w", err))
	}
	if err = _sql.Ping(); err != nil {
		panic(fmt.Errorf("failed ping: %s", err))
	}

	return engine
}
