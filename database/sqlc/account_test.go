package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ngochoangminh/golang_backend/util"
	"github.com/stretchr/testify/require"
)
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := tQuery.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)	
}

func TestGetAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	acc2, err := tQuery.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.Equal(t, acc1.Balance, acc2.Balance)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)	
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	arg := UpdateAccountParams {
		ID: acc1.ID,
		Balance: util.RandomMoney(),
	}
	acc2, err := tQuery.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc2)
	require.Equal(t, acc1.ID, acc2.ID)
	require.Equal(t, acc1.Currency, acc2.Currency)
	require.Equal(t, acc1.Owner, acc2.Owner)
	require.NotEqual(t, acc1.Balance, acc2.Balance)
	require.WithinDuration(t, acc1.CreatedAt, acc2.CreatedAt, time.Second)
}
func TestDeleteAccount(t *testing.T) {
	acc1 := createRandomAccount(t)
	err := tQuery.DeleteAccount(context.Background(), acc1.ID)
	require.NoError(t, err)

	acc2, err2 := tQuery.GetAccount(context.Background(), acc1.ID)
	require.Error(t, err2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())
	require.Empty(t, acc2)
}

func TestListAccount(t *testing.T) {
	// for i:=0;i<10;i++{
	// 	createRandomAccount(t)
	// }

	arg := ListAccountsParams{
		Owner: "ngoc",
		Limit: 5,
		Offset: 0,
	}

	accounts, err := tQuery.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 1)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}