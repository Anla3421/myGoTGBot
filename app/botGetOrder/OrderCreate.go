package botGetOrder

import (
	"fmt"
	"server/infrastructure/service/pbclient"
)

func OrderCreate() {
	res := pbclient.BotOrderCreate(pbclient.OrderCreateReq{})
	fmt.Println(res)
}
