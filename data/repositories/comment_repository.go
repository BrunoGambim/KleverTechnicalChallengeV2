package repositories

import entities "relembrando_2/domain/entities"

type CommentRepository interface {
	Insert(coment entities.Comment) error
	Update(comment entities.Comment) error
	FindById(id string) (entities.Comment, error)
	DeleteById(id string) error
	FindAll() ([]entities.Comment, error)
}
