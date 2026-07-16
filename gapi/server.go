package gapi

import (
	"fmt"

	db "github.com/Nevets9/simplebank/db/sqlc"
	"github.com/Nevets9/simplebank/token"
	"github.com/Nevets9/simplebank/util"
	pb "github.com/Nevets9/simplebank/pb"
)

// Server serves gRPC request for banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
	tokenMaker token.Maker
}

// NewGrpcServer create a new gRPC server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	//tokenMaker, err := token.NewPasetoMaker(config.TokenSecret)
	tokenMaker, err := token.NewJWTMaker(config.TokenSecret) 
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker :%s", err)
	}
	server := &Server{
		config: config,
		store: store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}