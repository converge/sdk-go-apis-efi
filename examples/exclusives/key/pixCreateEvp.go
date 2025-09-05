package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
	"github.com/converge/sdk-go-apis-efi/src/efipay/pix"
)

func main() {

	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{}{}

	res, err := efi.PixCreateEvp(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
