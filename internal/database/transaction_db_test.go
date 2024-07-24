package database_test

import (
	"database/sql"
	"testing"

	"github.com/ViniNepo/fc-ms-wallet-core/internal/database"
	"github.com/ViniNepo/fc-ms-wallet-core/internal/entity"
	"github.com/stretchr/testify/suite"

	_ "github.com/mattn/go-sqlite3"
)

type TransactionDBTestSuite struct {
	suite.Suite
	db            *sql.DB
	client        *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	TransactionDB *database.TransactionDB
}

func (t *TransactionDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	t.Nil(err)
	t.db = db

	db.Exec("CREATE table clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE table accounts (id varchar(255), client_id varchar(255), balance int, created_at date)")
	db.Exec("CREATE table transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date)")

	client, err := entity.NewClient("John", "j@j.com")
	t.Nil(err)
	t.client = client

	client2, err := entity.NewClient("John2", "jj@j.com")
	t.Nil(err)
	t.client2 = client2

	accountFrom := entity.NewAccount(client)
	accountFrom.Balance = 1000
	t.accountFrom = accountFrom

	accountTo := entity.NewAccount(client2)
	accountTo.Balance = 1000
	t.accountTo = accountTo

	t.TransactionDB = database.NewTransactionDB(db)
}

func (t *TransactionDBTestSuite) TearDownSuite() {
	t.db.Close()
	t.db.Exec("DROP table clients")
	t.db.Exec("DROP table accounts")
	t.db.Exec("DROP table transactions")
}

func TestClientDBTestSuit(t *testing.T) {
	suite.Run(t, new(TransactionDBTestSuite))
}

func (t *TransactionDBTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(t.accountFrom, t.accountTo, 100)
	t.Nil(err)

	err = t.TransactionDB.Create(transaction)
	t.Nil(err)
}
