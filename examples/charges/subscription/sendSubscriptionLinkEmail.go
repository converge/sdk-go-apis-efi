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
		"email": "gorbadoc.oldbuck@gerencianet.com.br",
	}

	res, err := efi.SendSubscriptionLinkEmail(1, body) // no lugar do 1 coloque o charge_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
