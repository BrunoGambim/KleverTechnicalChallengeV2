package network_repositories

import (
	network_errors "relembrando_2/data/network/errors"
	network_gateways "relembrando_2/data/network/gateways"
	"relembrando_2/domain/entities"
	"strconv"
)

type NetworkPostRepository struct {
	gateway *network_gateways.NetworkPostGateway
}

func NewNetworkPostRepository() *NetworkPostRepository {
	gateway := network_gateways.NewNetworkPostGateway()
	return &NetworkPostRepository{
		gateway: gateway,
	}
}

func (repository *NetworkPostRepository) FindById(id string) (entities.Post, error) {
	uintId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return entities.Post{}, network_errors.Handle(err)
	}

	datamapper, err := repository.gateway.GetPostById(int64(uintId))

	return datamapper.ToDomainEntity(), network_errors.Handle(err)
}
