package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	Port      uint64
	Interface string
	MysqlDsn  string
	RedisDsn  string
	CacheTTL  time.Duration
	Options   *LoadConfigOptions
	Logger    *slog.Logger
}

type LoadConfigOptions struct {
	PanicByErr bool
	RunMysql   bool
	RunRedis   bool
}

func LoadConfig(l ...LoadConfigOptions) *AppConfig {

	loadConfigOptions := &LoadConfigOptions{
		PanicByErr: true,
	}
	if len(l) > 0 {
		loadConfigOptions = &l[0]
	}
	portStr := os.Getenv("PORT")
	port, err := strconv.ParseUint(portStr, 10, 32)
	if err != nil {
		panic("bad config: PORT")
	}
	mysqlDsn := os.Getenv("MYSQL_DSN")
	redisDsn := os.Getenv("REDIS_DSN")
	cacheTTlStr := os.Getenv("CACHE_TTL")

	cacheTTl, err := strconv.Atoi(cacheTTlStr)
	if err != nil {
		panic("bad config: CACHE_TTL")
	}

	return &AppConfig{
		Port:      port,
		Interface: "",
		MysqlDsn:  mysqlDsn,
		RedisDsn:  redisDsn,
		Options:   loadConfigOptions,
		CacheTTL:  time.Duration(cacheTTl) * time.Second,
	}
}

func (c *AppConfig) GetServerAddress() string {
	return fmt.Sprintf("%s:%d", c.Interface, c.Port)
}

//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
