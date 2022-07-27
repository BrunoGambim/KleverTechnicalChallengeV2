package mongodb_repositories

import (
	"context"
	mongodb_data_mappers "relembrando_2/data/database/mongodb/data_mappers"
	mongodb_errors "relembrando_2/data/database/mongodb/errors"
	mongodb_gateways "relembrando_2/data/database/mongodb/gateways"
	"relembrando_2/data/repositories"
	entities "relembrando_2/domain/entities"
	"sync"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoCommentRepository struct {
	gateway        *mongodb_gateways.CommentGateway
	postRepository repositories.PostRepository
	sync.RWMutex
}

func NewMongoCommentRepository(ctx context.Context, postRepository repositories.PostRepository) (*MongoCommentRepository, error) {
	gateway, err := mongodb_gateways.NewCommentGateway(ctx)
	return &MongoCommentRepository{
		postRepository: postRepository,
		gateway:        gateway,
	}, mongodb_errors.Handle(err)
}

func (repository *MongoCommentRepository) Insert(coment entities.Comment) error {
	comment, err := mongodb_data_mappers.CommentFromDomainEntity(coment)

	if err != nil {
		return mongodb_errors.Handle(err)
	}

	repository.Lock()
	err = repository.gateway.Insert(comment)
	repository.Unlock()
	return mongodb_errors.Handle(err)
}

func (repository *MongoCommentRepository) Update(coment entities.Comment) error {
	comment, err := mongodb_data_mappers.CommentFromDomainEntity(coment)

	if err != nil {
		return mongodb_errors.Handle(err)
	}

	repository.Lock()
	err = repository.gateway.Update(comment)
	repository.Unlock()
	return mongodb_errors.Handle(err)
}

func (repository *MongoCommentRepository) FindById(id string) (entities.Comment, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return entities.Comment{}, mongodb_errors.Handle(err)
	}
	repository.RLock()
	commentDataMapper, err := repository.gateway.FindById(objectId)
	repository.RUnlock()

	if err != nil {
		return entities.Comment{}, mongodb_errors.Handle(err)
	}

	post, err := repository.postRepository.FindById(commentDataMapper.PostId)
	if err != nil {
		return entities.Comment{}, mongodb_errors.Handle(err)
	}

	return commentDataMapper.ToDomainEntity(post), mongodb_errors.Handle(err)
}

func (repository *MongoCommentRepository) DeleteById(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return mongodb_errors.Handle(err)
	}
	repository.Lock()
	err = repository.gateway.DeleteById(objectId)
	repository.Unlock()

	return mongodb_errors.Handle(err)
}

func (repository *MongoCommentRepository) FindAll() ([]entities.Comment, error) {
	repository.RLock()
	commentDataMapperList, err := repository.gateway.FindAll()
	repository.RUnlock()

	if err != nil {
		return []entities.Comment{}, mongodb_errors.Handle(err)
	}

	return repository.datamapperToList(commentDataMapperList)
}

func (repository *MongoCommentRepository) datamapperToList(commentDataMapperList []mongodb_data_mappers.CommentDataMapper) ([]entities.Comment, error) {
	commentList := []entities.Comment{}
	for _, commentDataMapper := range commentDataMapperList {
		post, err := repository.postRepository.FindById(commentDataMapper.PostId)
		if err != nil {
			return []entities.Comment{}, mongodb_errors.Handle(err)
		}

		commentList = append(commentList, commentDataMapper.ToDomainEntity(post))
	}
	return commentList, nil
}
