package gapi

import (
	"github.com/Nevets9/simplebank/pb"
	db "github.com/Nevets9/simplebank/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User{
	return &pb.User{
		Username: user.Username,
		FullName: user.FullName,
		Email: user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt) ,
		CreatedAt: timestamppb.New(user.CreatedAt) ,
	}
}