package botGetOrder

import (
	"fmt"
	"server/service/pbclient"
)

func OrderQuery(JWT string) (Message string) {
	res := pbclient.BotOrderQuery(JWT)

	storage := `OrderID : %v
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

	fmt.Println(Message)
	// fmt.Sprintf("OrderID :" + res.OrderID + "\n" +
	// 	"TotalPrice :" + res.TotalPrice + "\n" +
	// 	"Finish :" + res.Finish + "\n" +
	// 	"Account :" + res.OrderID + "\n" +
	// 	"Name :" + res.Name + "\n" +
	// 	"Amount :" + res.Amount + "\n" +
	// 	"Price :" + res.Price + "\n")

	return Message
}
