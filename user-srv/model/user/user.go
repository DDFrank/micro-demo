package user

import (
	"fmt"
	proto "micro-demo/user-srv/proto/user"
	"sync"
)

var (
	s *service
	m sync.RWMutex
)

type service struct {
}

// Service 用户服务类
type Service interface {
	QueryUserByName(userName string) (ret *proto.User, err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// 初始化用户服务层, 向model.go 暴露初始化方法
func Init() {
	m.Lock()
	defer m.Unlock()
	if s != nil {
		return
	}
	s = &service{}
}
