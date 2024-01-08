package repository

import (
	"context"
	appconfig "github.com/alichz2001/fidibo/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type AppRepo struct {
	db    *gorm.DB
	cache *redis.Client
}

func InitRepo(cfg *appconfig.AppConfig) *AppRepo {

	var mysql_ *gorm.DB
	var redis_ *redis.Client

	if cfg.Options.RunMysql {
		mysql_ = newMySqlRepo(cfg)
	}

	if cfg.Options.RunRedis {
		redis_ = newRedisRepo(cfg)
	}

	return &AppRepo{
		db:    mysql_,
		cache: redis_,
	}
}

func newMySqlRepo(cfg *appconfig.AppConfig) *gorm.DB {
	db, err := gorm.Open(mysql.Open(cfg.MysqlDsn), &gorm.Config{})
	PanicErr(err, cfg)
	return db
}

func newRedisRepo(cfg *appconfig.AppConfig) (client *redis.Client) {

	opt, err := redis.ParseURL(cfg.RedisDsn)
	PanicErr(err, cfg)

	client = redis.NewClient(opt)
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()

	status := client.Ping(ctx)
	PanicErr(status.Err(), cfg)

	return
}
