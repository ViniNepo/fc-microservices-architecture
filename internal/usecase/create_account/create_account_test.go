package create_account_test

import (
	"testing"

	"github.com/ViniNepo/fc-ms-wallet-core/internal/entity"
	"github.com/ViniNepo/fc-ms-wallet-core/internal/usecase/create_account"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ClientGatewayMock struct {
	mock.Mock
}

func (c *ClientGatewayMock) Get(id string) (*entity.Client, error) {
	args := c.Called(id)
	return args.Get(0).(*entity.Client), args.Error(1)
}

func (c *ClientGatewayMock) Save(client *entity.Client) error {
	args := c.Called(client)
	return args.Error(0)
}

type AccountGatewayMock struct {
	mock.Mock
}

func (a *AccountGatewayMock) Save(account *entity.Account) error {
	args := a.Called(account)
	return args.Error(0)
}

func (a *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
	args := a.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func TestCreateAccountUseCase_Execute(t *testing.T) {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	clientGatewayMock := &ClientGatewayMock{}
	clientGatewayMock.On("Get", client.ID).Return(client, nil)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("Save", mock.Anything).Return(nil)

	uc := create_account.NewCreateAccountUseCase(accountGatewayMock, clientGatewayMock)
	inputDTO := create_account.CreateAccountInputDTO{
		ClientID: client.ID,
	}

	outputDTO, err := uc.Execute(inputDTO)

	assert.NoError(t, err)
	assert.NotNil(t, outputDTO.ID)
	clientGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertExpectations(t)
	clientGatewayMock.AssertNumberOfCalls(t, "Get", 1)
	accountGatewayMock.AssertNumberOfCalls(t, "Save", 1)
}
