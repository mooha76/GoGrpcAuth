package service

import (
	"context"

	model "github.com/mooha76/GoGrpcAuth/model"
	"github.com/mooha76/GoGrpcAuth/pb"

	//"github.com/mooha76/GoGrpcAuth/token"
	usecase "github.com/mooha76/GoGrpcAuth/usecase/interface"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServiceServer struct {
	usecase usecase.AuthUseCase
	pb.UnimplementedAuthServiceServer
}

func NewAuthServiceServer(usecase usecase.AuthUseCase) pb.AuthServiceServer {
	return &authServiceServer{
		usecase: usecase,
	}
}

// User Sginup
func (c *authServiceServer) UserSignup(ctx context.Context, req *pb.UserSignupRequest) (*pb.UserSignupResponse, error) {

	signupRequest := model.UserSignupRequest{
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Age:       req.GetAge(),
		Email:     req.GetEmail(),
		Phone:     req.GetPhone(),
		Password:  req.GetPassword(),
	}

	UserId, err := c.usecase.UserSignup(ctx, signupRequest)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "%s", err.Error())
	}

	return &pb.UserSignupResponse{UserId: UserId}, nil
}
