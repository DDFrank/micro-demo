package handler

// handler调用model
import (
	"context"
	"log"
	user "micro-demo/user-srv/model/user"
	proto "micro-demo/user-srv/proto/user"
)
type Service struct{}

var (
	userService user.Service
)

// 初始化handler
func Init() {
	var err error
	userService, err = user.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误")
		return
	}
}

func (e *Service) QueryUserByName(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	user, err := userService.QueryUserByName(req.UserName)
	if err != nil {
		rsp.Success = false
		rsp.Error = &proto.Error{
			Code: 500,
			Detail: err.Error(),
		}
	}

	rsp.User = user
	rsp.Success = true
	return nil
}