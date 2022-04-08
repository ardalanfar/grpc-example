package client

import (
	"context"
	"fmt"
	"grpc-example/cmd"
	"grpc-example/conf"
	"io"
	"log"
	"strings"

	"google.golang.org/grpc"
)

func RunGrpcClient() {
	conn, err := grpc.Dial("127.0.0.1"+conf.GetConfig().Vars.Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := cmd.NewUserServiceClient(conn)

	input := ""
	fmt.Println("All People? (y/n)")
	fmt.Scanln(&input)

	//Get Users
	if strings.EqualFold(input, "y") {
		users, err := client.GetUsers(context.Background(), &cmd.Request{})
		if err != nil {
			log.Fatalln(err)
		}
		for {
			user, err := users.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(user)
		}
		return
	}

	//Get user by name
	fmt.Println("name?")
	fmt.Scanln(&input)
	user, err := client.GetUser(context.Background(), &cmd.Request{Name: input})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*user)
}
