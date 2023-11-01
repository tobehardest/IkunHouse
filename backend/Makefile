generateModel:
	generateVideoModel


generateVideo: generateVideoModel generateVideoProto


generateVideoModel:
	goctl model mysql ddl --src script/video.sql --dir cmd/video/model
	goctl model mysql ddl --src script/video_tag.sql --dir cmd/video/model
	goctl model mysql ddl --src script/tag.sql --dir cmd/video/model
	goctl model mysql ddl --src script/collect.sql --dir cmd/video/model

generateVideoProto:
	goctl rpc protoc cmd/video/rpc/video.proto --go_out=cmd/video/rpc --go-grpc_out=cmd/video/rpc --zrpc_out=cmd/video/rpc -style goZero