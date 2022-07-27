package post_service

import (
	"context"
	"relembrando_2/data/repositories"
	service_errors "relembrando_2/services/errors"
	"relembrando_2/services/post_service/post_assembler"
	"relembrando_2/services/post_service/post_pb"
)

type PostService struct {
	post_pb.UnimplementedPostServiceServer
	repository repositories.PostRepository
}

func NewPostService(repository repositories.PostRepository) *PostService {
	return &PostService{
		repository: repository,
	}
}

func (service *PostService) GetPostById(ctx context.Context, request *post_pb.PostIdRequest) (*post_pb.PostResponse, error) {
	id := post_assembler.IdRequestToString(request)
	post, err := service.repository.FindById(id)
	if err != nil {
		return nil, service_errors.Handle(err)
	}

	return post_assembler.PostEntityToResponse(post), service_errors.Handle(err)
}
