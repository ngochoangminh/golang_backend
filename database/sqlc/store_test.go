package database

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransection(t *testing.T) {
	store := NewStore(testDB)

	acc1 := createRandomAccount(t)
	acc2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)
	for i:= 0; i<n; i++ {
		go func () {
			res, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: acc1.ID,
				ToAccountID: acc2.ID,
				Amount: amount,
			})
			errs <- err
			results <- res
		}()
	}
	
	for i:=0; i<n; i++ {
		err := <- errs
		require.NoError(t, err)

		res := <- results
		require.NotEmpty(t, res)

		transfer := res.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, acc1.ID, transfer.FromAccountID)
		require.Equal(t, acc2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		entry := res.FromEntry
		require.NotEmpty(t, entry.Amount)


	}
	
}