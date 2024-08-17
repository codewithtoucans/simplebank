package api

import (
	"bytes"
	"encoding/json"
	"io"
	"testing"

	db "github.com/codewithtoucans/simplebank/db/sqlc"
	"github.com/codewithtoucans/simplebank/util"
	"github.com/stretchr/testify/require"
)

func TestGetAccoutAPI(t *testing.T) {
	// account := randomAccount()
	//
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	//
	// store := mockdb.NewMockStore(ctrl)
	//
	// store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Times(1).Return(account, nil)
	//
	// server := NewServer(store)
	// recorder := httptest.NewRecorder()
	//
	// url := fmt.Sprintf("/accounts/%d", account.ID)
	// request, err := http.NewRequest(http.MethodGet, url, nil)
	// require.NoError(t, err)
	//
	// server.router.ServeHTTP(recorder, request)
	//
	// require.Equal(t, http.StatusNotFound, recorder.Code)
	// requireBodyMatchAccount(t, recorder.Body, account)
}

func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account db.Account) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotAccount db.Account

	err = json.Unmarshal(data, &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)
}

func randomAccount(owner string) db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
