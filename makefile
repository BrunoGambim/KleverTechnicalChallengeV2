final: comment_service post_service upvote_service
comment_service: grpc_proto_files/comment.proto
	protoc --go_out=. --go_opt=module=relembrando_2 grpc_proto_files/comment.proto --go-grpc_out=. --go-grpc_opt=module=relembrando_2 grpc_proto_files/comment.proto
post_service: grpc_proto_files/post.proto
	protoc --go_out=./ --go_opt=module=relembrando_2 grpc_proto_files/post.proto --go-grpc_out=. --go-grpc_opt=module=relembrando_2 grpc_proto_files/post.proto
upvote_service: grpc_proto_files/upvote.proto
	protoc --go_out=./ --go_opt=module=relembrando_2 grpc_proto_files/upvote.proto --go-grpc_out=. --go-grpc_opt=module=relembrando_2 grpc_proto_files/upvote.proto
clean: 
	rm services/*/*.pb.go