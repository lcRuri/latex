#安装go-zero

go get -u github.com/zeromicro/go-zero

go install github.com/zeromicro/go-zero/tools/goctl@latest

goctl -v

#创建API服务

goctl api new core

cd core

#启动服务
go run core.go -f etc/core-api.yaml

go env -w GO111MODULE=on  #开启go mod管理，同理off就是关闭

http://localhost:8888/from/you

#使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero

#依赖
xorm.io/xorm

github.com/go-sql-driver/mysql

github.com/dgrijalva/jwt-go