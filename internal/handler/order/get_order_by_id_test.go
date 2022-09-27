package order

import (
	"context"
	"errors"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	mocks "order-mg/internal/mocks/service/order"
	"order-mg/internal/model"
	"testing"
	"time"
)

func TestGetOrderByID(t *testing.T) {
	type getOrderByID struct {
		mockID  int64
		mockOut model.Order
		mockErr error
	}
	type arg struct {
		givenID                string
		getOrderByID           getOrderByID
		getOrderByIDMockCalled bool
		expRs                  string
		expHTTPCode            int
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: "100",
			getOrderByID: getOrderByID{
				mockID: 100,
				mockOut: model.Order{
					Id:            100,
					Name:          "abc",
					Description:   "abc",
					TotalPrice:    99,
					Quantity:      10,
					Discount:      10,
					Shipping:      "abc",
					Status:        "APPROVED",
					UserId:        99,
					DatePurchased: time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
					CreatedAt:     time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
					UpdatedAt:     time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
				},
			},
			getOrderByIDMockCalled: true,
			expRs: `{
				"id":100,
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z",
				"created_at":"2022-09-21T16:30:00Z",
				"updated_at":"2022-09-21T16:30:00Z"
			}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: invalid id": {
			givenID:                "-100",
			getOrderByIDMockCalled: false,
			expRs:                  `{"code":"invalid_request", "description":"invalid id"}`,
			expHTTPCode:            http.StatusBadRequest,
		},
		"fail: error from service": {
			givenID: "100",
			getOrderByID: getOrderByID{
				mockID:  100,
				mockOut: model.Order{},
				mockErr: errors.New("something wrong"),
			},
			getOrderByIDMockCalled: true,
			expRs:                  `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode:            http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.OrderService)
			if tc.getOrderByIDMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetOrderByID", mock.Anything, tc.getOrderByID.mockID).
						Return(tc.getOrderByID.mockOut, tc.getOrderByID.mockErr),
				}
			}

			//GIVEN
			req := httptest.NewRequest(http.MethodGet, "/orders", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetOrderByID(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
