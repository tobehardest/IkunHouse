# api
goctl api go -api service/api-service/app.api -dir service/api-service -style goZero
# rpc
goctl rpc protoc service/user-service/rpc/user.proto --go_out=service/user-service/rpc/types --go-grpc_out=service/user-service/rpc/types --zrpc_out=service/user-service/rpc -style goZero
# model
goctl model mysql ddl --src service/user-service/model/user.sql --dir service/user-service/model