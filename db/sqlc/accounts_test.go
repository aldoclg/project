package db

import (
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/aldoclg/project/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Onwer:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Onwer, account.Onwer)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	a := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), a.ID)

	if err != nil {
		log.Fatal(err)
	}

	require.Equal(t, a.ID, account2.ID)

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Equal(t, account.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
}

func TestDeleteAccount(t *testing.T) {
	a := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), a.ID)

	require.NoError(t, err)

	a2, err := testQueries.GetAccount(context.Background(), a.ID)

	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, a2)

}
