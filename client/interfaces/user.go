package interfaces

import (
	"context"

	domain "github.com/mooha76/GoGrpcAuth/model"
)

type UserClient interface {
	SaveUser(ctx context.Context, user domain.UserSignupRequest) (userID uint64, err error)
}
