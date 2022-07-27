package repositories

import "relembrando_2/domain/entities"

type PostRepository interface {
	FindById(id string) (entities.Post, error)
}
