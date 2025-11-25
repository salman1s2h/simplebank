package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/salman1s2h/simplebank/db/mock"
	db "github.com/salman1s2h/simplebank/db/sqlc"
	"github.com/salman1s2h/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccountAPI(t *testing.T) {

	account := randomAccount()

	testCash := []struct {
		name          string
		AccountID     int64
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			AccountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountByID(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(account, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchAccount(t, recorder, account)
			},
		},
		{
			name:      "NotFound",
			AccountID: account.ID,
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetAccountByID(gomock.Any(), gomock.Eq(account.ID)).
					Times(1).
					Return(db.Account{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)

				data, err := io.ReadAll(recorder.Body)
				require.NoError(t, err)

				var response map[string]string
				err = json.Unmarshal(data, &response)
				require.NoError(t, err)

				require.Contains(t, response["error"], "not found")
			},
		},
	}
	for i := range testCash {
		tc := testCash[i]
		t.Run(tc.name, func(t *testing.T) {
			// build stubs
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			urlPath := fmt.Sprintf("/accounts/%d", account.ID)
			request, err := http.NewRequest(http.MethodGet, urlPath, nil)
			require.NoError(t, err)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)

		})

	}
}

func randomAccount() db.Account {
	// user := createRandomUser(nil)
	return db.Account{
		ID:    util.RandomInt(1, 1000),
		Owner: util.RandomOwner(),
		// Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
		// CreatedAt: time.Now(),
	}
}

func requireBodyMatchAccount(t *testing.T, body *httptest.ResponseRecorder, account db.Account) {
	data, err := io.ReadAll(body.Body)
	require.NoError(t, err)

	var gotAccount db.Account
	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	fmt.Println("Got Account:", gotAccount)
	fmt.Println("Expected Account:", account)
	require.Equal(t, account, gotAccount)
}
