package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
	"github.com/converge/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const inicio = "2021-03-01T03:01:35Z"
	const fim = "2021-03-05T22:01:35Z"

	res, err := efi.ListCharges(inicio, fim)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
