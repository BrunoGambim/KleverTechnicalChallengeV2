package upvote_assembler

import (
	"relembrando_2/domain/entities"
	"relembrando_2/services/upvote_service/upvote_pb"
)

func DeleteUpvoteRequestToCommentIdString(request *upvote_pb.DeleteUpvoteRequest) string {
	return request.CommentId
}

func DeleteUpvoteRequestToUpvoteIdString(request *upvote_pb.DeleteUpvoteRequest) string {
	return request.UpvoteId
}

func AddUpvoteRequestToIdString(request *upvote_pb.AddUpvoteRequest) string {
	return request.CommentId
}

func AddUpvoteRequestToEntity(request *upvote_pb.AddUpvoteRequest) entities.Upvote {
	return entities.NewUpvote(request.UpvoteType.String())
}

func UpvoteEntityListToResponse(upvotes []entities.Upvote) []*upvote_pb.UpvoteResponse {
	reponseList := []*upvote_pb.UpvoteResponse{}
	for _, upvote := range upvotes {
		reponseList = append(reponseList, UpvoteEntityToResponse(upvote))
	}

	return reponseList
}

func UpvoteEntityToResponse(upvote entities.Upvote) *upvote_pb.UpvoteResponse {
	return &upvote_pb.UpvoteResponse{
		Id:         upvote.GetId(),
		UpvoteType: upvote.GetUpvoteType(),
		CreatedAt:  upvote.GetFormatedCreatedAt(),
	}
}
