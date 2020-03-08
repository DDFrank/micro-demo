package user

import (
	"github.com/micro/go-micro/v2/util/log"
	"micro-demo/user-srv/basic/db"
	proto "micro-demo/user-srv/proto/user"
)
func (s *service) QueryUserByName(userName string) (ret *proto.User, err error) {
	queryString := `SELECT user_id, user_name, pwd  FROM user WHERE user_name = ?`

	o := db.GetDB()

	ret = &proto.User{}

	// 查询数据
	err = o.QueryRow(queryString, userName).Scan(&ret.Id, &ret.Name, &ret.Pwd)
	if err != nil {
		log.Logf("[QueryUserByName] 查询数据失败，err：%s", err)
		return
	}
	return
}
