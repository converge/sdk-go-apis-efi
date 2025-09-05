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
		"name": "My new plan",
	}

	res, err := efi.UpdatePlan(1, body) // no lugar do 1 coloque o plan_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
