package bankdb

import (
	"bankapi/db/util"
	"context"
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOwner(t *testing.T) {
	t.Log("NO EXISTING DETAILS FOUND")
	argsDetail := CreateOwnerDetailsParams{
		Name: util.RandomName(6),
		Age:  sql.NullInt32{Int32: util.RandomAge(10, 40), Valid: true},
	}

	details, _ := testQueries.GetOwner(context.Background(), argsDetail.Name)
	require.Empty(t, details)

	newOwnerDetails, err := testQueries.CreateOwnerDetails(context.Background(), argsDetail)
	if err != nil {
		log.Fatal(err)
	}
	id, err := newOwnerDetails.LastInsertId()
	t.Log("CREATED NEW DETAILS")
	t.Log("DETAILS ID NEWLY INSERTED", id)
	if err != nil {
		log.Fatal(err)
	}
	result, _ := testQueries.CreateOwner(context.Background(), int32(id))
	require.NotEmpty(t, result)

}
