package cache_service

import (
	"gin-web/pkg/global"
	"gin-web/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/piupuer/go-helper/pkg/query"
)

type RedisService struct {
	Q     query.Redis
	mysql service.MysqlService
}

func New(c *gin.Context) RedisService {
	ops := []func(*query.RedisOptions){
		query.WithRedisLogger(global.Log),
		query.WithRedisCtx(c),
		query.WithRedisDatabase(global.Conf.Mysql.DSN.DBName),
		query.WithRedisNamingStrategy(global.Mysql.NamingStrategy),
	}
	rd := RedisService{
		Q:     query.NewRedis(global.Redis, ops...),
		mysql: service.New(c),
	}
	return rd
}
