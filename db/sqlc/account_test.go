package db

import (
	"context"
	"testing"

	"github.com/codewithtoucans/simplebank/random"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func createRandomAccount(t *testing.T) Account {
	params := CreateAccountParams{
		Owner:    random.RandomOwner(),
		Balance:  random.RandomMoney(),
		Currency: random.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
