package mongodb_gateways

import (
	"context"
	"os"
	mongodb_database_factory "relembrando_2/data/database/mongodb"
	mongodb_data_mappers "relembrando_2/data/database/mongodb/data_mappers"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CommentGateway struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewCommentGateway(ctx context.Context) (*CommentGateway, error) {
	database, err := mongodb_database_factory.NewMongoDatabase(ctx)

	if err != nil {
		return nil, err
	}

	collectionName := os.Getenv("MONGODB_COMMENTS_COLLECTION")

	gateway := &CommentGateway{
		collection: database.GetCollection(collectionName),
		ctx:        ctx,
	}
	return gateway, err
}

func (gateway *CommentGateway) Insert(comment mongodb_data_mappers.CommentDataMapper) error {
	_, err := gateway.collection.InsertOne(gateway.ctx, comment)
	return err
}

func (gateway *CommentGateway) FindById(id primitive.ObjectID) (mongodb_data_mappers.CommentDataMapper, error) {
	result := mongodb_data_mappers.CommentDataMapper{}

	filter := bson.M{"_id": id}

	err := gateway.collection.FindOne(gateway.ctx, filter).Decode(&result)
	return result, err
}

func (gateway *CommentGateway) Update(comment mongodb_data_mappers.CommentDataMapper) error {
	filter := bson.M{"_id": comment.Id}
	update := bson.D{{Key: "$set", Value: comment}}

	_, err := gateway.collection.UpdateOne(gateway.ctx, filter, update)
	return err
}

func (gateway *CommentGateway) DeleteById(id primitive.ObjectID) error {
	filter := bson.M{"_id": id}

	_, err := gateway.collection.DeleteOne(gateway.ctx, filter)
	return err
}

func (gateway *CommentGateway) FindAll() ([]mongodb_data_mappers.CommentDataMapper, error) {
	filter := bson.M{}

	cur, err := gateway.collection.Find(gateway.ctx, filter)
	if err != nil {
		return nil, err
	}
	comments, err := curToCommentList(cur)
	return comments, err
}

func curToCommentList(cur *mongo.Cursor) ([]mongodb_data_mappers.CommentDataMapper, error) {
	comments := []mongodb_data_mappers.CommentDataMapper{}

	for cur.Next(context.TODO()) {
		comment := mongodb_data_mappers.CommentDataMapper{}
		err := cur.Decode(&comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
