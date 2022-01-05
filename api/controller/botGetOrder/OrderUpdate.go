package botGetOrder

import (
	"fmt"
	"server/service/pbclient"
)

func OrderUpdate(JWT string) {
	res := pbclient.BotOrderQuery(JWT)
	fmt.Println(res)
}
