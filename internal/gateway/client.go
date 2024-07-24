package gateway

import "github.com/ViniNepo/fc-ms-wallet-core/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(Client *entity.Client) error
}
