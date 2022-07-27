package comment_assembler

import (
	"relembrando_2/domain/entities"
	"relembrando_2/services/comment_service/comment_pb"
	post_assembler "relembrando_2/services/post_service/post_assembler"
	upvote_assembler "relembrando_2/services/upvote_service/upvote_assembler"
)

func IdRequestToString(request *comment_pb.IdRequest) string {
	return request.Id
}

func CommentRequestToEntity(request *comment_pb.CommentRequest, post entities.Post) entities.Comment {
	return entities.NewComment(request.Message, post)
}

func CommentRequestToIdString(request *comment_pb.CommentRequest) string {
	return request.PostId
}

func CommentEntityToResponse(comment entities.Comment) *comment_pb.CommentResponse {
	return &comment_pb.CommentResponse{
		Message: comment.GetMessage(),
		Id:      comment.GetId(),
		Upvotes: upvote_assembler.UpvoteEntityListToResponse(comment.GetUpvotes()),
		Post:    post_assembler.PostEntityToResponse(comment.GetPost()),
	}
}
