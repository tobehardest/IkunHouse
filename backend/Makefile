generateModel:
	generateVideoModel


generateVideo: generateVideoModel generateVideoProto


generateVideoModel:
	goctl model mysql ddl --src script/video.sql --dir cmd/video/model
	goctl model mysql ddl --src script/video_tag.sql --dir cmd/video/model
	goctl model mysql ddl --src script/tag.sql --dir cmd/video/model
	#goctl model mysql ddl --src script/collect.sql --dir cmd/video/model

generatePostProto:
	goctl rpc protoc cmd/post/rpc/post.proto --go_out=cmd/post/rpc --go-grpc_out=cmd/post/rpc --zrpc_out=cmd/post/rpc -style goZero
generatePostModel:
	goctl model mysql ddl --src script/post.sql --dir cmd/post/model
	goctl model mysql ddl --src script/post_tag.sql --dir cmd/post/model
	goctl model mysql ddl --src script/tag.sql --dir cmd/post/model

generateVideoProto:
	goctl rpc protoc cmd/video/rpc/video.proto --go_out=cmd/video/rpc --go-grpc_out=cmd/video/rpc --zrpc_out=cmd/video/rpc -style goZero