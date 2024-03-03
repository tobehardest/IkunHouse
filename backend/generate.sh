
#1.根据api文件生成api目录handler、logic
    cd backend/cmd/api/desc
    goctl api go --api api.api --dir ../ -style goZero

# 2.根据proto文件生成rpc目录对应的client、server、internal
# 以 video.proto 文件生成为例
    cd backend/cmd/video
    goctl rpc protoc video.proto --go_out=. --go-grpc_out=. --zrpc_out=. -style goZero

# model
goctl model mysql ddl --src service/user-service/model/user.sql --dir service/user-service/model

# 4. 根据api文件生成swagger json文件
# 4.1 首先需要安装goctl-swagger插件
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest
# 4.2 进入到cmd/api/desc 目录下
goctl api plugin -plugin goctl-swagger="swagger -filename api.json" -api api.api -dir .
# 该命令会在当前目录下生成api.json文件