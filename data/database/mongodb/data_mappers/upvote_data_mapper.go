package mongodb_data_mappers

import (
	"relembrando_2/domain/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpvoteDataMapper struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	UpvoteType string             `bson:"upvote_type"`
	CreatedAt  int64              `bson:"created_at"`
}

func UpvoteFromEntity(upvote entities.Upvote) (UpvoteDataMapper, error) {
	var id primitive.ObjectID
	var err error
	if upvote.GetId() != "" {
		id, err = primitive.ObjectIDFromHex(upvote.GetId())
	} else {
		id = primitive.NewObjectIDFromTimestamp(time.Now())
	}

	if err != nil {
		return UpvoteDataMapper{}, err
	}

	return UpvoteDataMapper{
		Id:         id,
		UpvoteType: upvote.GetUpvoteType(),
		CreatedAt:  upvote.GetCreatedAt(),
	}, nil
}

func (upvoteDataMapper *UpvoteDataMapper) ToDomainEntity() entities.Upvote {
	return entities.LoadUpvote(upvoteDataMapper.Id.Hex(), upvoteDataMapper.UpvoteType, upvoteDataMapper.CreatedAt)
}

func ToUpvoteList(dataMapperList []UpvoteDataMapper) []entities.Upvote {
	entityList := []entities.Upvote{}
	for _, dataMapper := range dataMapperList {
		entityList = append(entityList, dataMapper.ToDomainEntity())
	}
	return entityList
}

func FromUpvoteEntityList(entityList []entities.Upvote) ([]UpvoteDataMapper, error) {
	dataMapperList := []UpvoteDataMapper{}
	for _, entity := range entityList {
		dataMapper, err := UpvoteFromEntity(entity)
		if err != nil {
			return nil, err
		}
		dataMapperList = append(dataMapperList, dataMapper)
	}
	return dataMapperList, nil
}
