package interfaces

import (
	"context"

	model "github.com/mooha76/GoGrpcAuth/model"
)

type AuthUseCase interface {
	UserSignup(ctx context.Context, user model.UserSignupRequest) (CreatedUser string, err error)
}
