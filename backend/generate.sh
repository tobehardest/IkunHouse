
#1.根据api文件生成api目录handler、logic
    cd backend/cmd/api/desc
    goctl api go --api api.api --dir ../ -style goZero

# 2.根据proto文件生成rpc目录对应的client、server、internal
# 以 video.proto 文件生成为例
    cd backend/cmd/video
    goctl rpc protoc video.proto --go_out=. --go-grpc_out=. --zrpc_out=. -style goZero

# model
goctl model mysql ddl --src script/user.sql --dir cmd/user/model