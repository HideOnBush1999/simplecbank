package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store defines all functions to interact with the database
type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

// SQLStore 连接到真实的数据库，并实现 Store 接口
// SQLStore defines all functions to execute SQL queries and transactions
type SQLStore struct {
	// Queries 只能执行一条语句，所以不能执行事务
	// 现在采用组合的方式来拓展结构体的功能
	*Queries // 结构体指针 （Store 会得到 Queries 的方法）
	db       *sql.DB
}

// 为什么要这样组合呢？
// Queries 包含 CURD 函数，*sql.DB类型的 db 变量包含和 transaction 相关的函数

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}


