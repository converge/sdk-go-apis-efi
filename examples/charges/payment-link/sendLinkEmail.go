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
		"email": "oldbuck@gerencianet.com.br",
	}

	res, err := efi.SendLinkEmail(1, body) //no lugar do 1 informe o id da cobrança desejada

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
