package gateway

import "github.com/ViniNepo/fc-ms-wallet-core/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(ID string) (*entity.Account, error)
}
