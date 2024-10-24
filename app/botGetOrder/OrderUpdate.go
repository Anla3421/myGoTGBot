package botGetOrder

import (
	"fmt"
	"server/infrastructure/service/pbclient"
)

func OrderUpdate(JWT string) {
	res := pbclient.BotOrderQuery(JWT)
	fmt.Println(res)
}
