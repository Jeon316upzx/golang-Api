package bankdb

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateOwner(t *testing.T) {
	argsDetail := CreateOwnerDetailsParams{
		Name: "Ifeanyi",
		Age:  sql.NullInt32{Int32: 20, Valid: true},
	}

	details, err := testQueries.GetOwner(context.Background(), argsDetail.Name)
	fmt.Println(details)
	t.Log(details)
	require.NoError(t, err)
	require.NotEmpty(t, details)
	if err != nil {
		log.Fatal(err)
	}

	t.Log(details)

	// result, err := testQueries.CreateOwner(context.Background())

}
