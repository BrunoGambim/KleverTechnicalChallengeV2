package upvote_service

import (
	"context"

	repositories "relembrando_2/data/repositories"
	service_errors "relembrando_2/services/errors"
	upvote_assembler "relembrando_2/services/upvote_service/upvote_assembler"
	upvote_pb "relembrando_2/services/upvote_service/upvote_pb"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type UpvoteService struct {
	upvote_pb.UnimplementedUpvoteServiceServer
	commentRepository repositories.CommentRepository
}

func NewUpvoteService(commentRepository repositories.CommentRepository) *UpvoteService {
	return &UpvoteService{
		commentRepository: commentRepository,
	}
}

func (service *UpvoteService) AddUpvoteToComment(ctx context.Context, request *upvote_pb.AddUpvoteRequest) (*empty.Empty, error) {
	comment, err := service.commentRepository.FindById(upvote_assembler.AddUpvoteRequestToIdString(request))
	if err != nil {
		return &empty.Empty{}, service_errors.Handle(err)
	}
	upvote := upvote_assembler.AddUpvoteRequestToEntity(request)
	comment.AddUpvote(upvote)

	err = service.commentRepository.Update(comment)
	return &empty.Empty{}, service_errors.Handle(err)
}

func (service *UpvoteService) DeleteUpvoteFromComment(ctx context.Context, request *upvote_pb.DeleteUpvoteRequest) (*empty.Empty, error) {
	comment, err := service.commentRepository.FindById(upvote_assembler.DeleteUpvoteRequestToCommentIdString(request))
	if err != nil {
		return &empty.Empty{}, service_errors.Handle(err)
	}
	comment.RemoveUpvote(upvote_assembler.DeleteUpvoteRequestToUpvoteIdString(request))

	err = service.commentRepository.Update(comment)
	return &empty.Empty{}, service_errors.Handle(err)
}
