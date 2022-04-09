package server

import (
	"context"
	"grpc-example/cmd"
	"grpc-example/database"
)

type Grpcserver struct {
	Dbhandler *database.SQLhandler
}

func NewGServer() *Grpcserver {
	return &Grpcserver{
		Dbhandler: database.CreateDBConnection(),
	}
}

//Get user by name and send
func (server *Grpcserver) GetUser(ctx context.Context, r *cmd.Request) (*cmd.User, error) {
	user, err := server.Dbhandler.GetUserByName(r.GetName())
	if err != nil {
		return nil, err
	}
	grpcUser := convertToGrpcUser(user)
	return grpcUser, nil
}

//Get users and send by Stream
func (server *Grpcserver) GetUsers(r *cmd.Request, stream cmd.UserService_GetUsersServer) error {
	users, err := server.Dbhandler.GetUsres()
	if err != nil {
		return err
	}

	for _, user := range users {
		grpcUsers := convertToGrpcUser(user)
		err := stream.Send(grpcUsers)
		if err != nil {
			return err
		}
	}
	return nil
}

//Convert to gRps user struct
func convertToGrpcUser(user database.User) *cmd.User {
	return &cmd.User{
		Id:     int32(user.ID),
		Name:   user.Name,
		Family: user.Family,
	}
}
