module micro-demo/user-srv

go 1.13

require (
	github.com/DDFrank/micro-demo/basic v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.3.4
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.2.0
	github.com/micro/protoc-gen-micro/v2 v2.0.0 // indirect
)

replace github.com/DDFrank/micro-demo/basic => ../basic
