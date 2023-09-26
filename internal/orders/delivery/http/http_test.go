package http

import (
	"bytes"

	"github.com/ilhame90/gymshark-task/internal/config"
	"github.com/ilhame90/gymshark-task/internal/domain/mocks"
	"github.com/ilhame90/gymshark-task/internal/models"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func TestPostAccounts(t *testing.T) {

	testCases := []struct {
		name          string
		body          []byte
		buildStubs    func(usecase *mocks.MockUsecase)
		checkResponse func(recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: []byte(`{"ordered_items": 1}`),
			buildStubs: func(uc *mocks.MockUsecase) {
				uc.EXPECT().NumberOfPacks(gomock.Any(), gomock.Eq(1)).Times(1).Return([]models.Pack{{Name: 250, Quantity: 1}})
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				//requireBodyMatchAccount(t, recorder.Body, respAccount)
			},
		},
		{
			name: "Invalid order",
			body: []byte(`"ordered_items":-1`),
			buildStubs: func(uc *mocks.MockUsecase) {
				uc.EXPECT().NumberOfPacks(gomock.Any(), gomock.Any()).Times(0)
			},
			checkResponse: func(recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			uc := mocks.NewMockUsecase(ctrl)
			tc.buildStubs(uc)
			rec := httptest.NewRecorder()

			e := echo.New()
			grp := e.Group("/v1")
			cfg := &config.Config{}
			handler := NewOrdersHandler(cfg, uc)
			RegisterHandlers(grp, handler)

			req := httptest.NewRequest(http.MethodPost, "/v1/order", bytes.NewReader(tc.body))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			e.ServeHTTP(rec, req)
			tc.checkResponse(rec)
		})
	}
}

// func requireBodyMatchAccount(t *testing.T, body *bytes.Buffer, account Account) {
// 	data, err := ioutil.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotAccount Account
// 	err = json.Unmarshal(data, &gotAccount)
// 	require.NoError(t, err)
// 	require.Equal(t, account.Id, gotAccount.Id)
// 	require.Equal(t, account.Currency, gotAccount.Currency)
// 	require.Equal(t, account.Balance, gotAccount.Balance)
// }

// func requireBodyMatchBalance(t *testing.T, body *bytes.Buffer, balance BalanceResponse) {
// 	data, err := ioutil.ReadAll(body)
// 	require.NoError(t, err)

// 	var gotBalance BalanceResponse
// 	err = json.Unmarshal(data, &gotBalance)
// 	require.NoError(t, err)
// 	require.Equal(t, balance, gotBalance)
// }
