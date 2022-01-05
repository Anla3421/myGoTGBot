package botGetOrder

import (
	"fmt"
	"server/service/pbclient"
)

func OrderQuery(JWT string) (Message string) {
	res := pbclient.BotOrderQuery(JWT)

	storage := `你的訂單資訊如下：
	系統訂單編號 : %v
	總價格 : %v
	完成繳費? : %s
	購物者帳號 : %v
	商品名稱 : %s
	商品數量 : %v
	商品價格 : %v
	`

	demo := `你的訂單資訊如下：
	OrderID : %v
	TotalPrice : %v
	Finish : %s
	Account : %v
	Name : %s
	Amount : %v
	Price : %v
	`
	Message = fmt.Sprintf(storage,
		res.OrderID,
		res.TotalPrice,
		res.Finish,
		res.OrderID,
		res.Name,
		res.Amount,
		res.Price)

	DemoMessage := fmt.Sprintf(demo,
		res.OrderID,
		res.TotalPrice,
		res.Finish,
		res.OrderID,
		res.Name,
		res.Amount,
		res.Price)

	fmt.Println(DemoMessage)

	return Message
}
