package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
	"github.com/converge/sdk-go-apis-efi/src/efipay"
)

func main() {

	credentials := configs.Credentials
	efi := efipay.NewEfiPay(credentials)

	customer := map[string]interface{}{
		"name":         "Gorbadoc Oldbuck",
		"cpf":          "04267484171",
		"phone_number": "5144916523",
	}

	items := []map[string]interface{}{
		{
			"name":   "Item 1",
			"value":  1000,
			"amount": 1,
		},
		{
			"name":   "Item 2",
			"value":  2000,
			"amount": 2,
		},
	}

	body := map[string]interface{}{
		"items":       items,
		"customer":    customer,
		"expire_at":   "2020-12-02",
		"repeats":     5,
		"split_items": false,
	}

	res, err := efi.CreateCarnet(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
