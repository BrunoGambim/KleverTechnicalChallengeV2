package dependence_providers

import (
	"context"
	"log"
	mongodb_repositories "relembrando_2/data/database/mongodb/repositories"
	network_repositories "relembrando_2/data/network/repositories"
	"relembrando_2/data/repositories"
)

func GetCommentRepository(ctx context.Context) repositories.CommentRepository {
	repository, err := mongodb_repositories.NewMongoCommentRepository(ctx, GetPostRepository())
	if err != nil {
		log.Fatal(err)
	}
	return repository
}

func GetPostRepository() repositories.PostRepository {
	repository := network_repositories.NewNetworkPostRepository()
	return repository
}
