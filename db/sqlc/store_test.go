package bankdb

import (
	"bankapi/db/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	argsDetail := CreateOwnerDetailsParams{
		Name: "Jeon",
		Age:  sql.NullInt32{Int32: util.RandomAge(10, 40), Valid: true},
	}

	details, _ := testQueries.GetOwner(context.Background(), argsDetail.Name)

	id := details.ID
	argsAccount := CreateAccountParams{
		Balance:  int64(util.RandomAmount(0, 20000)),
		Currency: util.RandomCurrency(),
		Owner:    sql.NullInt32{Int32: id, Valid: true},
	}

	createdAcc, _ := testQueries.CreateAccount(context.Background(), argsAccount)

	argsDetail1 := CreateOwnerDetailsParams{
		Name: "Jeon",
		Age:  sql.NullInt32{Int32: util.RandomAge(10, 40), Valid: true},
	}

	details1, _ := testQueries.GetOwner(context.Background(), argsDetail1.Name)

	id1 := details1.ID
	argsAccount1 := CreateAccountParams{
		Balance:  int64(util.RandomAmount(0, 20000)),
		Currency: util.RandomCurrency(),
		Owner:    sql.NullInt32{Int32: id1, Valid: true},
	}

	createdAcc1, _ := testQueries.CreateAccount(context.Background(), argsAccount1)

	accountID1, _ := createdAcc.LastInsertId()
	accountID2, _ := createdAcc1.LastInsertId()
	amount := int64(10)
	result, err := store.TransferTx(context.Background(), TransferTxParams{
		FromAccountID: int32(accountID1),
		ToAccountID:   int32(accountID2),
		Amount:        amount,
	})

	require.NoError(t, err)
	require.NotEmpty(t, result)
}
