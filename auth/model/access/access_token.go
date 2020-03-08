package access

import (
	"fmt"
	"github.com/DDFrank/micro-demo/basic/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/micro/go-micro/v2/broker"
	log "github.com/micro/go-micro/v2/logger"
	"time"
)

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"

	tokenExpiredTopic = "mu.micro.book.topic.auth.tokenExpired"
)

// token 的持有者
type Subject struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

// 生成token并保存到redis
func (s *service) MakeAccessToken(subject *Subject) (ret string, err error) {
	m, err := s.createTokenClaims(subject)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token Claim 失败，err: %s", err)
	}

	// 创建
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)
	ret, err = token.SignedString([]byte(config.GetJwtConfig().GetSecretKey()))
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 创建token失败，err: %s", err)
	}

	// 保存到redis
	err = s.saveTokenToCache(subject, ret)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}

	return
}

// 获取token
func (s *service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}
	return
}

// 清除用户token
func (s *service) DelUserAccessToken(tokenString string) (err error) {
	// 解析token字符串
	claims, err := s.parseToken(tokenString)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
	}

	// 通过解析到的用户id删除
	err = s.delTokenFromCache(&Subject{
		ID: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
	}

	// 广播删除
	msg := &broker.Message{
		Body: []byte(claims.Subject),
	}

	if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
		log.Infof("[pub] 发布token删除消息失败： %v", err)
	} else {
		fmt.Println("[pub] 发布token删除消息：", string(msg.Body))
	}

	return
}
