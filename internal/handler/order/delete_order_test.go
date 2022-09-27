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
	"testing"
)

func TestDeleteOrder(t *testing.T) {
	type deleteOrder struct {
		mockIn  int64
		mockErr error
	}
	type arg struct {
		givenID               string
		deleteOrder           deleteOrder
		deleteOrderMockCalled bool
		expRs                 string
		expHTTPCode           int
	}
	tcs := map[string]arg{
		"success: ": {
			givenID: "123",
			deleteOrder: deleteOrder{
				mockIn: 123,
			},
			deleteOrderMockCalled: true,
			expHTTPCode:           http.StatusOK,
			expRs:                 `{"message":"Deleted Order"}`,
		},
		"fail: invalid ID": {
			givenID:               "abc",
			deleteOrderMockCalled: false,
			expHTTPCode:           http.StatusBadRequest,
			expRs:                 `{"code":"invalid_request", "description":"id must be a number"}`,
		},
		"fail: error from service": {
			givenID: "123",
			deleteOrder: deleteOrder{
				mockIn:  123,
				mockErr: errors.New("something wrong"),
			},
			deleteOrderMockCalled: true,
			expHTTPCode:           http.StatusInternalServerError,
			expRs:                 `{"code":"internal_server_error", "description":"Something went wrong please try again later!"}`,
		},
	}
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//MOCK
			mockSvc := new(mocks.OrderService)
			if tc.deleteOrderMockCalled {
				mockSvc.ExpectedCalls = []*mock.Call{
					mockSvc.On("DeleteOrder", mock.Anything, tc.deleteOrder.mockIn).
						Return(tc.deleteOrder.mockErr),
				}
			}

			//GIVEN
			req := httptest.NewRequest(http.MethodDelete, "/orders", nil)
			routeCtx := chi.NewRouteContext()
			routeCtx.URLParams.Add("id", tc.givenID)

			ctx := context.WithValue(req.Context(), chi.RouteCtxKey, routeCtx)
			req = req.WithContext(ctx)
			res := httptest.NewRecorder()

			//WHEN
			instance := New(mockSvc)
			instance.DeleteOrder(res, req)

			//THEN
			require.Equal(t, tc.expHTTPCode, res.Code)
			require.JSONEq(t, tc.expRs, res.Body.String())
			mockSvc.AssertExpectations(t)
		})
	}
}
