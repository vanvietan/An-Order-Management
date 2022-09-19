package user

import (
	"context"
	"net/http"
	"testing"
	"time"

	mocks "order-mg/internal/mocks/service/user"
	"order-mg/internal/model"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetUserByIDTest(t *testing.T) {
	// httpmock.Activate()
	// defer httpmock.DeactivateAndReset()

	path := "http://localhost:3000/users/"

	type arg struct {
		givenID      int64
		givenRequest string
		mockOutput   model.Users
		expRs        model.Users
		mockErr      error
		expErr       error
		expHTTPCode  int
	}

	tcs := map[string]arg{
		"success: ": {
			givenID:      101,
			givenRequest: path + "101",
			mockOutput: model.Users{
				Id:          101,
				Name:        "abc",
				Username:    "abc1",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "ADMIN",
				CreatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
			},
			expRs: model.Users{
				Id:          101,
				Name:        "abc",
				Username:    "abc1",
				PhoneNumber: "123",
				Address:     "abc",
				Age:         1,
				Role:        "ADMIN",
				CreatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
				UpdatedAt:   time.Date(2022, 3, 15, 16, 0, 0, 0, time.UTC),
			},
			expHTTPCode: http.StatusOK,
		},
	}
	ctx := context.Background()
	for s, tc := range tcs {
		t.Run(s, func(t *testing.T) {
			//GIVEN
			instance := new(mocks.UserService)
			instance.On("GetUserByID", mock.Anything, tc.givenID).Return(tc.mockOutput, tc.mockErr)

			//WHEN
			h := New(instance)
			rs, err := h.UserSvc.GetUserByID(ctx, tc.givenID)

			//THEN
			if tc.expErr != nil {
				require.EqualError(t, err, tc.expErr.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expRs, rs)
			}

		})
	}
}
