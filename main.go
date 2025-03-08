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

	fmt.Println("Customer deleted successfully")
	customer, err := payarc.GetCustomer("DPpKjVKDAMKxVnMN ")
	if err != nil {
		fmt.Println("Error getting customer:", err)
		return
	}
	fmt.Println("Customer details:", customer)
}
