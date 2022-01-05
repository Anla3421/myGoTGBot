package pbclient

import (
	"context"
	"fmt"
	"log"

	// protobuf "github.com/Anla3421/myGoProtobuf/myGoMemberServer/go"
	protobuf "server/service/pbclient/protobuf"

	"google.golang.org/grpc"
)

//gPRC client 連線建立
func init() {
	fmt.Println("gRPC client initial")
	CreateConn()
}

var Client protobuf.BotGetOrderClient

func CreateConn() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	Client = protobuf.NewBotGetOrderClient(conn)
}

//gPRC client function
func BotOrderQuery(JWT string) *protobuf.BotOrderQueryResponse {
	req := &protobuf.BotOrderQueryRequest{
		JWT: JWT,
	}
	res, err := Client.BotOrderQuery(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC client:error while calling BotOrderQuery Service: %v \n", err)
	}
	log.Printf("gRPC client:Response from BotOrderQuery Service: %v", res)
	return res
}

func BotOrderUpdate(JWT string) *protobuf.BotOrderUpdateResponse {
	req := &protobuf.BotOrderUpdateRequest{
		JWT: JWT,
	}
	res, err := Client.BotOrderUpdate(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC client:error while calling BotOrderUpdate Service: %v \n", err)
	}
	log.Printf("gRPC client:Response from BotOrderUpdate Service: %v", res)
	return res
}
