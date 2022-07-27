package network_data_mappers

import (
	"relembrando_2/domain/entities"
	"strconv"
)

type PostDataMapper struct {
	Id     uint64 `json:"id"`
	UserId uint64 `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func PostFromDomainEntity(post entities.Post) (PostDataMapper, error) {
	var id uint64
	var userId uint64
	var err error

	if post.GetId() != "" {
		id, err = strconv.ParseUint(post.GetId(), 10, 64)
	}
	if err != nil {
		return PostDataMapper{}, err
	}

	userId, err = strconv.ParseUint(post.GetUserId(), 10, 64)
	if err != nil {
		return PostDataMapper{}, err
	}

	return PostDataMapper{
		Id:     id,
		UserId: userId,
		Title:  post.GetTitle(),
		Body:   post.GetBody(),
	}, nil
}

func (postDataMapper *PostDataMapper) ToDomainEntity() entities.Post {
	return entities.LoadPost(strconv.Itoa(int(postDataMapper.Id)), strconv.Itoa(int(postDataMapper.UserId)), postDataMapper.Title, postDataMapper.Body)
}
