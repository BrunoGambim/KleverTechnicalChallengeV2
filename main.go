package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"relembrando_2/dependence_providers"
	"relembrando_2/services/comment_service"
	"relembrando_2/services/comment_service/comment_pb"
	"relembrando_2/services/post_service"
	"relembrando_2/services/post_service/post_pb"
	"relembrando_2/services/upvote_service"
	"relembrando_2/services/upvote_service/upvote_pb"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func loadEnvFile() {
	err := godotenv.Overload(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func getNetListener() net.Listener {
	port := os.Getenv("PORT")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}

func main() {
	loadEnvFile()
	ctx := context.Background()

	list := getNetListener()
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	comment_pb.RegisterCommentServiceServer(grpcServer,
		comment_service.NewCommentService(dependence_providers.GetCommentRepository(ctx), dependence_providers.GetPostRepository()))

	post_pb.RegisterPostServiceServer(grpcServer,
		post_service.NewPostService(dependence_providers.GetPostRepository()))

	upvote_pb.RegisterUpvoteServiceServer(grpcServer,
		upvote_service.NewUpvoteService(dependence_providers.GetCommentRepository(ctx)))

	err := grpcServer.Serve(list)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
