package interfaces

import (
	"context"

	model "github.com/mooha76/GoGrpcAuth/model"
)

type UserClient interface {
	FindUserByEmail(ctx context.Context, email string) (model.User, error)
	FindUserByPhone(ctx context.Context, phone string) (model.User, error)
	SaveUser(ctx context.Context, user model.UserSignupRequest) (userID uint64, err error)
}
