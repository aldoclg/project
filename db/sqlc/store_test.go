package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < 5; i++ {
		txtName := fmt.Sprintf("tx %d", i)
		go func() {
			ctx := context.WithValue(context.Background(), txKey, txtName)
			result, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result

		}()
	}

	for i := 0; i < 5; i++ {
		require.NoError(t, <-errs)

		result := <-results
		require.NotEmpty(t, result)
		require.NotEmpty(t, result.Transfer)
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.Equal(t, amount, result.Transfer.Amount)

		_, err := store.GetTransfer(context.Background(), result.Transfer.ID)
		require.NoError(t, err)
	}

}
