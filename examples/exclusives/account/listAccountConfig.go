package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/src/efipay/pix"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)


	res, err := efi.ListAccountConfig(nil) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
