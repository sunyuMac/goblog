package bootstrap

import (
	"goblog/app/models/article"
	"goblog/app/models/category"
	"goblog/app/models/user"
	"goblog/pkg/config"
	"goblog/pkg/logger"
	"goblog/pkg/model"
	"goblog/pkg/redis"
	"gorm.io/gorm"
	"time"
)

// SetupDB 初始化DB
func SetupDB() {
	setupMysql()
	setupRedis()
}

// setupMysql 初始化数据库和 ORM
func setupMysql() {
	// 建立数据库连接池
	db := model.ConnectDB()

	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(config.GetInt("database.mysql.max_open_connections"))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(config.GetInt("database.mysql.max_idle_connections"))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.mysql.max_life_seconds")) * time.Second)

	// 创建和维护数据表结构
	migration(db)
}

// migration 维护数据库表结构
func migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&user.User{},
		&article.Article{},
		&category.Category{},
	)
	logger.LogError(err)
}

// setupRedis 初始化redis
func setupRedis() {
	redis.ConnectRedisPool()
}
