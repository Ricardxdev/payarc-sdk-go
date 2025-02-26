package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	payarcsdk "github.com/Ricardxdev/payarc-sdk-go/pkg"
)

func main() {
	payarc := payarcsdk.NewPayarcClient(context.Background(), payarcsdk.PayarcClientOptions{
		BaseUrl:      os.Getenv("PAYARC_URL"),
		Version:      "1.2.0",
		ApiVersion:   "v1",
		PayarcPrefix: "payarc",
		Token:        os.Getenv("PAYARC_TOKEN"),
		HTTPClient:   &http.Client{Timeout: 2 * time.Minute},
	})

	//payarc.DeleteCard("jAPDKVpNjMp4VNxn", "P1Lv0Nmm9vP205MN")

	customers, err := payarc.GetCustomers(1, 10)
	if err != nil {
		panic(err)
	}

	for _, customer := range customers.Data {
		if len(customer.Card.Data) != 0 {
			fmt.Println(customer.CustomerID)
			// _, err := payarc.UpdateCard("P1Lv0Nmm9vP205MN", payarcin.UpdateCardDTO{
			// 	CardHolderName: "Ricardo Martinez",
			// 	ExpMonth:       12,
			// 	ExpYear:        2025,
			// })

			// if err != nil {
			// 	panic(err)
			// }

			fmt.Println(">>>")
			cards, err := payarc.GetCustomerCards(customer.CustomerID)
			if err != nil {
				panic(err)
			}
			for _, card := range cards.Cards {
				fmt.Println(card.ID, card.IsDefault)
			}
			break
		}
	}
}
