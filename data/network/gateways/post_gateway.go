package network_gateways

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	clients "relembrando_2/data/network/clients"
	data_mappers "relembrando_2/data/network/data_mappers"
	network_errors "relembrando_2/data/network/errors"
	"strconv"
)

type NetworkPostGateway struct {
	client  clients.PostApiClient
	fullUrl string
}

func NewNetworkPostGateway() *NetworkPostGateway {
	postClient := clients.NewPostApiClient()
	path := os.Getenv("POST_API_POST_PATH")
	fullUrl := fmt.Sprintf("%s/%s", postClient.BaseUrl, path)
	return &NetworkPostGateway{
		client:  *postClient,
		fullUrl: fullUrl,
	}
}

func (gateway *NetworkPostGateway) GetPostById(id int64) (data_mappers.PostDataMapper, error) {
	var post data_mappers.PostDataMapper

	response, err := gateway.client.Get(fmt.Sprintf("%s/%s", gateway.fullUrl, strconv.Itoa(int(id))))
	if err != nil {
		return data_mappers.PostDataMapper{}, err
	}
	defer response.Body.Close()

	err = network_errors.HandleHttpStatus(response.StatusCode)
	if err != nil {
		return data_mappers.PostDataMapper{}, err
	}

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return data_mappers.PostDataMapper{}, err
	}

	err = json.Unmarshal(content, &post)

	return post, err
}
