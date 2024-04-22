package db

import (
	"context"

	"github.com/gucooing/hkrpg-go/dispatch/config"
	"github.com/gucooing/hkrpg-go/pkg/alg"
	"github.com/gucooing/hkrpg-go/pkg/logger"
)

var ctx = context.Background()

// NewStore 创建一个新的 store。
func NewStore(config *config.Config) *Store {
	s := &Store{config: config}
	accountMysqlConf := config.MysqlConf["account"]
	s.AccountMysql = alg.NewMysql(accountMysqlConf.Dsn)
	s.AccountMysql.AutoMigrate(&Account{})

	redisLoginConf := config.RedisConf["player_login"]
	s.LoginRedis = alg.NewRedis(redisLoginConf.Addr, redisLoginConf.Password, redisLoginConf.DB)

	logger.Info("数据库连接成功")
	return s
}
