package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
	"github.com/converge/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	body := map[string]interface{}{
		"custom_id":        "Product 0001",
		"notification_url": "http://domain.com/notification",
	}

	res, err := efi.UpdateSubscriptionMetadata(1, body) // no lugar do 1 coloque o subscription_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
