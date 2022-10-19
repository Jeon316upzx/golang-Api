package bankdb

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOwnerDetails(t *testing.T) {
	arg := CreateOwnerDetailsParams{
		Name: "Ifeanyi",
		Age:  sql.NullInt32{Int32: 20, Valid: true},
	}

	result, err := testQueries.CreateOwnerDetails(context.Background(), arg)
	require.NoError(t, err)
	fmt.Println(result, "QUERY RESULT")
	require.NotEmpty(t, result)
}
