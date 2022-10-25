package bankdb

import (
	"bankapi/db/util"
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOwnerDetails(t *testing.T) {
	arg := CreateOwnerDetailsParams{
		Name: util.RandomName(6),
		Age:  sql.NullInt32{Int32: util.RandomAge(10, 40), Valid: true},
	}

	result, err := testQueries.CreateOwnerDetails(context.Background(), arg)
	require.NoError(t, err)
	fmt.Println(result, "QUERY RESULT")
	require.NotEmpty(t, result)
}
