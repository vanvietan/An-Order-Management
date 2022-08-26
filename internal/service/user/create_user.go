package user

import (
	"context"
	"fmt"
	"order-mg/internal/model"

	"github.com/sony/sonyflake"
	"golang.org/x/crypto/bcrypt"
)

var sf *sonyflake.Sonyflake

const (
	MinCost     int = 4
	MaxCost     int = 31
	DefaultCost int = 10
)

// CreateUser create a user
func (i impl) CreateUser(ctx context.Context, user model.Users) (model.Users, error) {
	// id, err := sf.NextID()
	// if err != nil {
	// 	fmt.Errorf("error when get a user id, %v", user.Id)
	// 	return model.Users{}, err
	// }

	// user.Id = sonyflakeId(id)
	user.Password = hashPassword(user.Password)
	_, errs := i.userRepo.CreateUser(ctx, user)
	if errs != nil {
		fmt.Errorf("error when get a user, %v", user.Name)
		return model.Users{}, errs
	}
	return user, nil

}

func hashPassword(s string) string {
	bs, err := bcrypt.GenerateFromPassword([]byte(s), MinCost)
	if err != nil {
		fmt.Errorf("error when hash pw, %v", bs)
		return ""
	}
	bpass := string(bs)
	return bpass
}

// func sonyflakeId(id uint64) int64 {
// 	body, err := json.Marshal(sonyflake.Decompose(id))
// 	if err != nil {
// 		fmt.Errorf("error when hash pw, %v", body)
// 		return 0
// 	}
// 	sf, _ := strconv.ParseInt(string(body), 10, 64)

// 	return sf
// }
