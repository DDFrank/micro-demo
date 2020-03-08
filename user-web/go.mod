module micro-demo/user-web

go 1.13

require (
	github.com/DDFrank/micro-demo/auth v0.0.0-00010101000000-000000000000
	github.com/DDFrank/micro-demo/basic v0.0.0-00010101000000-000000000000
	github.com/DDFrank/micro-demo/user-srv v0.0.0-00010101000000-000000000000
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.2.0
)

replace github.com/DDFrank/micro-demo/user-srv => ../user-srv

replace github.com/DDFrank/micro-demo/auth => ../auth

replace github.com/DDFrank/micro-demo/basic => ../basic
