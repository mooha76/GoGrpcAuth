package usecase

import (
	"context"
	"fmt"
	"time"

	client "github.com/mooha76/GoGrpcAuth/client/interface"
	domain "github.com/mooha76/GoGrpcAuth/model"

	interfaces "github.com/mooha76/GoGrpcAuth/usecase/interface"
	"github.com/mooha76/GoGrpcAuth/utils"
)

type authUsecase struct {
	userClient client.UserClient
}

const (
	accessTokenDuration  = time.Minute * 20
	refreshTokenDuration = time.Hour * 24 * 7
)

func NewAuthUsecase(userClient client.UserClient) interfaces.AuthUseCase {
	return &authUsecase{

		userClient: userClient,
	}
}

// Signup for user
func (c *authUsecase) UserSignup(ctx context.Context, user domain.UserSignupRequest) (CreatedUser string, err error) {

	existUser, err := c.userClient.FindUserByEmail(ctx, user.Email)
	if err != nil {
		return "", fmt.Errorf("failed to check user already exist \nerror:%s", err.Error())
	}
	// if user exist and verified also then return error
	if existUser.ID != 0 && existUser.Verified {
		return "", fmt.Errorf("user already exist with given details and verified")
	}

	userID := existUser.ID
	// if user not exist then save the user
	if userID == 0 {

		user.Password, err = utils.GenerateHashFromPassword(user.Password)
		if err != nil {
			return "", fmt.Errorf("failed to hash user password")
		}

		userID, err = c.userClient.SaveUser(ctx, user)
		if err != nil {
			return "", fmt.Errorf("failed to save user \nerror%s", err.Error())
		}
	}

	return CreatedUser, nil
}
