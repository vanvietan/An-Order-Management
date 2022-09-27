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
	"strings"
	"testing"
	"time"
)

func TestUpdateOrder(t *testing.T) {
	type updateOrder struct {
		mockID  int64
		mockIn  model.Order
		mockOut model.Order
		mockErr error
	}
	type arg struct {
		givenID               string
		givenBody             string
		updateOrder           updateOrder
		updateOrderMockCalled bool
		expRs                 string
		expHTTPCode           int
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: "100",
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrder: updateOrder{
				mockID: 100,
				mockIn: model.Order{
					Name:          "abc",
					Description:   "abc",
					TotalPrice:    99,
					Quantity:      10,
					Discount:      10,
					Shipping:      "abc",
					Status:        "APPROVED",
					UserId:        99,
					DatePurchased: time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
				},
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
			updateOrderMockCalled: true,
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
		"fail: invalid Body invalid name": {
			givenBody: `{
				"name":"",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid name"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid description": {
			givenBody: `{
				"name":"abc",
				"description":"",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid description"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid total price": {
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price": 0,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid price"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid quantity": {
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":0,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid quantity"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid discount": {
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":-10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid discount"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid shipping": {
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid shipping method"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid status": {
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"GLORY",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid status"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: invalid Body invalid userID": {
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":-99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid userID"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: error invalid ID": {
			givenID: "0",
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID":99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: false,
			expRs:                 `{"code":"invalid_request", "description":"invalid id"}`,
			expHTTPCode:           http.StatusBadRequest,
		},
		"fail: error from service ": {
			givenID: "100",
			givenBody: `{
				"name":"abc",
				"description":"abc",
				"total_price":99,
				"quantity":10,
				"discount":10,
				"shipping":"abc",
				"status":"APPROVED",
				"userID": 99,
				"date_purchased":"2022-09-21T16:30:00Z"
			}`,
			updateOrderMockCalled: true,
			updateOrder: updateOrder{
				mockID: 100,
				mockIn: model.Order{
					Name:          "abc",
					Description:   "abc",
					TotalPrice:    99,
					Quantity:      10,
					Discount:      10,
					Shipping:      "abc",
					Status:        "APPROVED",
					UserId:        99,
					DatePurchased: time.Date(2022, 9, 21, 16, 30, 0, 0, time.UTC),
				},
				mockOut: model.Order{},
				mockErr: errors.New("something wrong"),
			},
			expRs:       `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
			expHTTPCode: http.StatusInternalServerError,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.OrderService)
			if tc.updateOrderMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("UpdateOrder", mock.Anything, tc.updateOrder.mockIn, tc.updateOrder.mockID).
						Return(tc.updateOrder.mockOut, tc.updateOrder.mockErr),
				}
			}

			//GIVEN
			req := httptest.NewRequest(http.MethodPut, "/orders", strings.NewReader(tc.givenBody))
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.UpdateOrder(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
