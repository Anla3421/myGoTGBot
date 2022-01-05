package pbclient

import (
	"context"
	"log"
	protobuf "server/service/pbclient/protobuf"
)

type OrderCreateReq struct {
	ItemID     []int64
	Amount     []int64
	Price      []int64
	JWT        string
	TotalPrice int32
}

func BotOrderCreate(Req OrderCreateReq) *protobuf.BotOrderCreateResponse {
	req := &protobuf.BotOrderCreateRequest{
		JWT: Req.JWT,
	}
	res, err := Client.BotOrderCreate(context.Background(), req)
	if err != nil {
		log.Fatalf("gRPC client:error while calling BotOrderCreate Service: %v \n", err)
	}
	log.Printf("gRPC client:Response from BotOrderCreate Service: %v", res)
	return res
}
