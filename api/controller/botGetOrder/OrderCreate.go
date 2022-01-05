package botGetOrder

import (
	"fmt"
	"server/service/pbclient"
)

func OrderCreate() {

	res := pbclient.BotOrderCreate(pbclient.OrderCreateReq{})
	fmt.Println(res)
}
