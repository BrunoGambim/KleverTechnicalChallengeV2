package mongodb_data_mappers

import (
	"relembrando_2/domain/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentDataMapper struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	PostId  string             `bson:"postId"`
	Message string             `bson:"message"`
	Upvotes []UpvoteDataMapper `bson:"upvotes"`
}

func CommentFromDomainEntity(comment entities.Comment) (CommentDataMapper, error) {
	var id primitive.ObjectID
	var err error
	if comment.GetId() != "" {
		id, err = primitive.ObjectIDFromHex(comment.GetId())
	}

	if err != nil {
		return CommentDataMapper{}, err
	}

	upvotes, err := FromUpvoteEntityList(comment.GetUpvotes())
	if err != nil {
		return CommentDataMapper{}, err
	}

	post := comment.GetPost()

	return CommentDataMapper{
		Id:      id,
		PostId:  post.GetId(),
		Message: comment.GetMessage(),
		Upvotes: upvotes,
	}, nil
}

func (commentDataMapper *CommentDataMapper) ToDomainEntity(post entities.Post) entities.Comment {
	return entities.LoadComment(commentDataMapper.Id.Hex(), post, commentDataMapper.Message, ToUpvoteList(commentDataMapper.Upvotes))
}
