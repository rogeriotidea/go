package gateway

import "github.com/devfullcycle/fc-ms-wallet/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error 
	FindById(id string) (*entity.Account, error)
}