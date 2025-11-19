package db

import (
	"context"
	"testing"
	"time"

	"github.com/salman1s2h/simplebank/util"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	arg := CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: "Secrect",
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user.Username, user2.Username)
	require.Equal(t, user.HashedPassword, user.HashedPassword)
	require.Equal(t, user.FullName, user.FullName)
	require.Equal(t, user.Email, user.Email)
	require.WithinDuration(t, user.CreatedAt, user.CreatedAt, time.Second)
}

// func TestUpdateAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)

// 	arg := UpdateAccountParams{
// 		ID:      account1.ID,
// 		Balance: util.RandomMoney(),
// 	}
// 	account2, err := testQueries.UpdateAccount(context.Background(), arg)
// 	// account2, err := testQueries.UpdateAccount(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, account2)

// 	require.Equal(t, account1.ID, account2.ID)
// 	require.Equal(t, account1.Owner, account2.Owner)
// 	require.Equal(t, arg.Balance+account1.Balance, account2.Balance)
// 	require.Equal(t, account1.Currency, account2.Currency)
// 	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
// }

// func TestDeleteAccount(t *testing.T) {
// 	account1 := createRandomAccount(t)
// 	err := testQueries.DeleteAccount(context.Background(), account1.ID)
// 	require.NoError(t, err)

// 	account2, err := testQueries.GetAccountByID(context.Background(), account1.ID)
// 	require.Error(t, err)
// 	require.EqualError(t, err, sql.ErrNoRows.Error())
// 	require.Empty(t, account2)
// }

// func TestListAccounts(t *testing.T) {
// 	var lastAccount Account
// 	for i := 0; i < 10; i++ {
// 		lastAccount = createRandomAccount(t)
// 	}

// 	arg := GetAccountsParams{
// 		// Owner:  lastAccount.Owner,
// 		Limit:  5,
// 		Offset: 0,
// 	}

// 	accounts, err := testQueries.GetAccounts(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, accounts)

// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 		require.Equal(t, lastAccount.Owner, account.Owner)
// 	}
// }

// func TestListAccounts(t *testing.T) {
// 	// var lastAccount Account
// 	// for i := 0; i < 10; i++ {
// 	// 	lastAccount = createRandomAccount(t)
// 	// }

// 	// Build parameters for the query
// 	arg := GetAccountsParams{
// 		Limit:  5,
// 		Offset: 0,
// 	}

// 	// Call the function
// 	accounts, err := testQueries.GetAccounts(context.Background(), arg)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, accounts)

// 	// Validate the returned data
// 	for _, account := range accounts {
// 		require.NotEmpty(t, account)
// 		require.NotZero(t, account.ID)
// 		require.NotEmpty(t, account.Currency)
// 		require.NotZero(t, account.CreatedAt)
// 	}
// }
