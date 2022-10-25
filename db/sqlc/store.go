package bankdb

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err %v , rb err %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxResult struct {
	Transfer        Transfer    `json:"transfer"`
	FromAccount     Account     `json: "from_account"`
	ToAccount       Account     `json: "to_account"`
	FromTransaction Transaction `json: "from_transaction"`
	ToTransaction   Transaction `json: "to_transaction"`
}

type TransferTxParams struct {
	FromAccountID int32 `json: "from_account_id"`
	ToAccountID   int32 `json: "to_account_id"`
	Amount        int64 `json: "amount"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		res, er := q.CreateTransfer(ctx, CreateTransferParams{
			Amount:      arg.Amount,
			ToAccount:   arg.ToAccountID,
			FromAccount: arg.FromAccountID,
		})
		err = er
		id, _ := res.LastInsertId()
		ttt, _ := q.GetTransfer(ctx, int32(id))
		result.Transfer = ttt

		if err != nil {
			return err
		}

		res1, err1 := q.CreateTransaction(ctx, CreateTransactionParams{
			Type:    TransactionsTypeDebit,
			Amount:  -arg.Amount,
			Account: arg.FromAccountID,
		})

		err = err1
		id1, _ := res1.LastInsertId()
		ttt1, _ := q.GetTransaction(ctx, int32(id1))
		result.FromTransaction = ttt1

		if err != nil {
			return err
		}

		res2, err2 := q.CreateTransaction(ctx, CreateTransactionParams{
			Type:    TransactionsTypeCredit,
			Amount:  +arg.Amount,
			Account: arg.ToAccountID,
		})

		err = err2
		id2, _ := res2.LastInsertId()
		ttt2, _ := q.GetTransaction(ctx, int32(id2))
		result.ToTransaction = ttt2

		if err != nil {
			return err
		}

		return nil

	})
	return result, err
}
