package post_assembler

import (
	"relembrando_2/domain/entities"
	"relembrando_2/services/post_service/post_pb"
)

func IdRequestToString(request *post_pb.PostIdRequest) string {
	return request.Id
}

func PostEntityToResponse(post entities.Post) *post_pb.PostResponse {
	return &post_pb.PostResponse{
		Id:     post.GetId(),
		UserId: post.GetUserId(),
		Title:  post.GetTitle(),
		Body:   post.GetBody(),
	}
}
