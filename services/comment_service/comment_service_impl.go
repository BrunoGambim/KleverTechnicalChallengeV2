package comment_service

import (
	"context"

	repositories "relembrando_2/data/repositories"
	comment_assembler "relembrando_2/services/comment_service/comment_assembler"
	comment_pb "relembrando_2/services/comment_service/comment_pb"
	service_errors "relembrando_2/services/errors"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type CommentService struct {
	comment_pb.UnimplementedCommentServiceServer
	commentRepository repositories.CommentRepository
	postRepository    repositories.PostRepository
}

func NewCommentService(commentRepository repositories.CommentRepository,
	postRepository repositories.PostRepository) *CommentService {
	return &CommentService{
		commentRepository: commentRepository,
		postRepository:    postRepository,
	}
}

func (service *CommentService) GetCommentById(ctx context.Context, request *comment_pb.IdRequest) (*comment_pb.CommentResponse, error) {
	id := comment_assembler.IdRequestToString(request)
	comment, err := service.commentRepository.FindById(id)
	return comment_assembler.CommentEntityToResponse(comment), service_errors.Handle(err)
}

func (service *CommentService) DeleteCommentById(ctx context.Context, request *comment_pb.IdRequest) (*empty.Empty, error) {
	id := comment_assembler.IdRequestToString(request)
	err := service.commentRepository.DeleteById(id)
	return &empty.Empty{}, service_errors.Handle(err)
}

func (service *CommentService) GetAllComments(empty *empty.Empty, stream comment_pb.CommentService_GetAllCommentsServer) error {
	commentList, err := service.commentRepository.FindAll()

	if err != nil {
		return service_errors.Handle(err)
	}

	for _, comment := range commentList {
		response := comment_assembler.CommentEntityToResponse(comment)
		err = stream.Send(response)
		if err != nil {
			return service_errors.Handle(err)
		}
	}

	return service_errors.Handle(err)
}

func (service *CommentService) InsertComment(ctx context.Context, request *comment_pb.CommentRequest) (*empty.Empty, error) {
	post, err := service.postRepository.FindById(comment_assembler.CommentRequestToIdString(request))
	if err != nil {
		return &empty.Empty{}, service_errors.Handle(err)
	}
	comment := comment_assembler.CommentRequestToEntity(request, post)
	err = service.commentRepository.Insert(comment)
	return &empty.Empty{}, service_errors.Handle(err)
}
