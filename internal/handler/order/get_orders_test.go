package order

import (
	"context"
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

func TestGetOrders(t *testing.T) {
	type getOrders struct {
		mockCursor int64
		mockLimit  int
		mockOut    []model.Order
		mockErr    error
	}
	type arg struct {
		givenCursor         string
		givenLimit          string
		getOrders           getOrders
		getOrdersMockCalled bool
		expRs               string
		expHTTPCode         int
	}
	tcs := map[string]arg{
		"success: all orders": {
			givenCursor: "0",
			givenLimit:  "20",
			getOrders: getOrders{
				mockCursor: 0,
				mockLimit:  20,
				mockOut: []model.Order{
					{
						Id:            101,
						Name:          "abc1",
						Description:   "abc1",
						TotalPrice:    101,
						Quantity:      101,
						Discount:      11,
						Shipping:      "abc",
						Status:        model.StatusApproved,
						UserId:        101,
						DatePurchased: time.Date(2022, 3, 16, 0, 0, 0, 0, time.UTC),
						CreatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
						UpdatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
					},
					{
						Id:            100,
						Name:          "abc",
						Description:   "abc",
						TotalPrice:    100,
						Quantity:      100,
						Discount:      10,
						Shipping:      "abc",
						Status:        model.StatusApproved,
						UserId:        100,
						DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
						CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
						UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
					},
					{
						Id:            99,
						Name:          "an",
						Description:   "an",
						TotalPrice:    99,
						Quantity:      99,
						Discount:      9,
						Shipping:      "hcm",
						Status:        model.StatusShipped,
						UserId:        99,
						DatePurchased: time.Date(2022, 3, 14, 0, 0, 0, 0, time.UTC),
						CreatedAt:     time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
						UpdatedAt:     time.Date(2022, 3, 14, 14, 0, 0, 0, time.UTC),
					},
				},
			},
			getOrdersMockCalled: true,
			expRs: `{
					"orders": [
						{
							"id": 101,
							"name": "abc1",
							"description": "abc1",
							"total_price": 101,
							"quantity": 101,
							"discount": 11,
							"shipping": "abc",
							"status": "APPROVED",
							"userID": 101,
							"date_purchased": "2022-03-16T00:00:00Z",
							"created_at": "2022-03-16T16:00:00Z",
							"updated_at": "2022-03-16T16:00:00Z",
							"deleted_at": null
						},
						{
							"id": 100,
							"name": "abc",
							"description": "abc",
							"total_price": 100,
							"quantity": 100,
							"discount": 10,
							"shipping": "abc",
							"status": "APPROVED",
							"userID": 100,
							"date_purchased": "2022-03-15T00:00:00Z",
							"created_at": "2022-03-15T15:00:00Z",
							"updated_at": "2022-03-15T15:00:00Z",
							"deleted_at": null
						},
						{
							"id": 99,
							"name": "an",
							"description": "an",
							"total_price": 99,
							"quantity": 99,
							"discount": 9,
							"shipping": "hcm",
							"status": "SHIPPED",
							"userID": 99,
							"date_purchased": "2022-03-14T00:00:00Z",
							"created_at": "2022-03-14T14:00:00Z",
							"updated_at": "2022-03-14T14:00:00Z",
							"deleted_at": null
						}
					],
					"cursor": 99
					}`,
			expHTTPCode: http.StatusOK,
		},
		"success: limit 2": {
			givenCursor: "0",
			givenLimit:  "2",
			getOrders: getOrders{
				mockCursor: 0,
				mockLimit:  2,
				mockOut: []model.Order{
					{
						Id:            101,
						Name:          "abc1",
						Description:   "abc1",
						TotalPrice:    101,
						Quantity:      101,
						Discount:      11,
						Shipping:      "abc",
						Status:        model.StatusApproved,
						UserId:        101,
						DatePurchased: time.Date(2022, 3, 16, 0, 0, 0, 0, time.UTC),
						CreatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
						UpdatedAt:     time.Date(2022, 3, 16, 16, 0, 0, 0, time.UTC),
					},
					{
						Id:            100,
						Name:          "abc",
						Description:   "abc",
						TotalPrice:    100,
						Quantity:      100,
						Discount:      10,
						Shipping:      "abc",
						Status:        model.StatusApproved,
						UserId:        100,
						DatePurchased: time.Date(2022, 3, 15, 0, 0, 0, 0, time.UTC),
						CreatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
						UpdatedAt:     time.Date(2022, 3, 15, 15, 0, 0, 0, time.UTC),
					},
				},
			},
			getOrdersMockCalled: true,
			expRs: `{
						"orders": [
						{
							"id": 101,
							"name": "abc1",
							"description": "abc1",
							"total_price": 101,
							"quantity": 101,
							"discount": 11,
							"shipping": "abc",
							"status": "APPROVED",
							"userID": 101,
							"date_purchased": "2022-03-16T00:00:00Z",
							"created_at": "2022-03-16T16:00:00Z",
							"updated_at": "2022-03-16T16:00:00Z",
							"deleted_at": null
						},
						{
							"id": 100,
							"name": "abc",
							"description": "abc",
							"total_price": 100,
							"quantity": 100,
							"discount": 10,
							"shipping": "abc",
							"status": "APPROVED",
							"userID": 100,
							"date_purchased": "2022-03-15T00:00:00Z",
							"created_at": "2022-03-15T15:00:00Z",
							"updated_at": "2022-03-15T15:00:00Z",
							"deleted_at": null
						}
					],
						"cursor": 100
					}`,
			expHTTPCode: http.StatusOK,
		},
		"fail: invalid limit": {
			givenLimit:          "abc",
			givenCursor:         "0",
			getOrdersMockCalled: false,
			expRs:               `{"code":"invalid_request", "description":"limit must be a number"}`,
			expHTTPCode:         http.StatusBadRequest,
		},
		"fail: invalid lastID": {
			givenLimit:          "20",
			givenCursor:         "abc",
			getOrdersMockCalled: false,
			expRs:               `{"code":"invalid_request", "description":"cursor must be a number"}`,
			expHTTPCode:         http.StatusBadRequest,
		},
		"success: empty": {
			givenLimit:          "3",
			givenCursor:         "1",
			getOrdersMockCalled: true,
			getOrders: getOrders{
				mockLimit:  3,
				mockCursor: 1,
				mockOut:    []model.Order{},
				mockErr:    nil,
			},
			expRs:       `{"orders":null,"cursor":0}`,
			expHTTPCode: http.StatusOK,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.OrderService)
			if tc.getOrdersMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("GetOrders", mock.Anything, tc.getOrders.mockLimit, tc.getOrders.mockCursor).
						Return(tc.getOrders.mockOut, tc.getOrders.mockErr),
				}
			}
			//GIVEN
			path := "/orders" + "?limit=" + tc.givenLimit + "&cursor=" + tc.givenCursor
			req := httptest.NewRequest(http.MethodGet, path, nil)
			routeCtx := chi.NewRouteContext()
			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.GetOrders(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}

}
