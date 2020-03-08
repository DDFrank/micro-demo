package access

import (
	"fmt"
	"github.com/DDFrank/micro-demo/basic/redis"
	r "github.com/go-redis/redis"
	"sync"
)

var (
	s           *service
	redisClient *r.Client
	m           sync.RWMutex
)

type service struct{}

// Service 用户服务
type Service interface {
	// 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// 清除用户token
	DelUserAccessToken(token string) (err error)
}

// 获取服务
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	redisClient = redis.GetRedis()

	s = &service{}
}
