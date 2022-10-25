package bankdb

import (
	"bankapi/db/util"
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

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

	createdAcc, err := testQueries.CreateAccount(context.Background(), argsAccount)

	if err != nil {
		log.Fatal(err)
	}

	lastID, _ := createdAcc.LastInsertId()
	require.NotZero(t, lastID)

}

func TestGetAccount(t *testing.T) {

	argsDetail := CreateOwnerDetailsParams{
		Name: "Jeon",
		Age:  sql.NullInt32{Int32: util.RandomAge(10, 40), Valid: true},
	}

	details, _ := testQueries.GetOwner(context.Background(), argsDetail.Name)

	account, _ := testQueries.GetAccount(context.Background(), sql.NullInt32{Int32: details.ID, Valid: true})

	require.NotEmpty(t, account)
}

func TestUpdateAccount(t *testing.T) {
	argsDetail := CreateOwnerDetailsParams{
		Name: "Jeon",
		Age:  sql.NullInt32{Int32: util.RandomAge(10, 40), Valid: true},
	}

	details, _ := testQueries.GetOwner(context.Background(), argsDetail.Name)

	profile, _ := testQueries.GetProfile(context.Background(), details.ID)
	t.Log(profile.ID, "PROFILE")

	updateBal := util.RandomAmount(0, 20000)

	id := profile.ID
	argsUpdate := UpdateAccountParams{
		Balance: int64(updateBal),
		Owner:   sql.NullInt32{Int32: id, Valid: true},
	}

	updatedAccount, _ := testQueries.UpdateAccount(context.Background(), argsUpdate)

	// LASTID, _ := updatedAccount.LastInsertId()
	// fetchedAccount, _ := testQueries.GetAccountById(context.Background(), sql.NullInt32{Int32: int32(LASTID), Valid: true})

	require.NotEmpty(t, updatedAccount)

}
