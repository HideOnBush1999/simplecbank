package db

import "context"

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

// 在创建新用户的同一个数据库事务中，将任务发送到 redis
// 将任务发送到 redis 是通过 AfterCreate 这个回调函数实现的
func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err!= nil {
			return err
		}

		return arg.AfterCreate(result.User)
	})

	return result, err
}
