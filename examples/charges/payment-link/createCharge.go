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
		"items": []map[string]interface{}{
			{
				"name":   "Product 1",
				"value":  1000,
				"amount": 2,
			},
		},
		"shippings": []map[string]interface{}{
			{
				"name":  "Default Shipping Cost",
				"value": 100,
			},
		},
	}

	res, err := efi.CreateCharge(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
