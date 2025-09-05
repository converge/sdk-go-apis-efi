package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
	"github.com/converge/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	const txid = "adssshdsjdsjeyccdyddsasdstxid01"

	res, err := efi.DetailDueCharge(txid)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
