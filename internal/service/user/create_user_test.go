package user

// import (
// 	"context"
// 	mocks "order-mg/internal/mocks/repository/user"
// 	"order-mg/internal/model"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// )

// func TestCreateUser(t *testing.T) {
// 	type createUser struct {
// 		mockCreate model.Users
// 		mockResp   model.Users
// 		mockErr    error
// 	}
// 	type arg struct {
// 		createUser createUser
// 		expRs      model.Users
// 		expErr     error
// 	}

// 	tcs := map[string]arg{
// 		"success: create a user ": {
// 			createUser: createUser{
// 				mockCreate: model.Users{
// 					Name:        "an",
// 					Username:    "an",
// 					Password:    "abc",
// 					PhoneNumber: "123",
// 					Address:     "abc",
// 					Age:         1,
// 					Role:        "USER",
// 					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
// 					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
// 				},
// 				mockResp: model.Users{
// 					Id:          103,
// 					Name:        "an",
// 					Username:    "an",
// 					Password:    "abc",
// 					PhoneNumber: "123",
// 					Address:     "abc",
// 					Age:         1,
// 					Role:        "USER",
// 					CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
// 					UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
// 				},
// 			},
// 			expRs: model.Users{
// 				Id:          103,
// 				Name:        "an",
// 				Username:    "an",
// 				PhoneNumber: "123",
// 				Address:     "abc",
// 				Age:         1,
// 				Role:        "USER",
// 				CreatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
// 				UpdatedAt:   time.Date(2022, 4, 15, 16, 0, 0, 0, time.UTC),
// 			},
// 		},
// 	}

// 	ctx := context.Background()

// 	for s, tc := range tcs {
// 		t.Run(s, func(t *testing.T) {
// 			//GIVEN
// 			instance := new(mocks.UserRepository)
// 			instance.On("CreateUser", ctx, tc.createUser.mockCreate).Return(tc.createUser.mockResp, tc.createUser.mockErr)

// 			//WHEN
// 			svc := New(instance)
// 			rs, err := svc.CreateUser(ctx, tc.createUser.mockCreate)

// 			//THEN
// 			if err != nil {
// 				require.EqualError(t, err, tc.expErr.Error())
// 			} else {
// 				require.Equal(t, tc.expRs, rs)
// 			}

// 		})
// 	}
// }
