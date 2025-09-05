package main

import (
	"fmt"
	"github.com/converge/sdk-go-apis-efi/src/efipay/pix"
	"github.com/converge/sdk-go-apis-efi/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	efi := pix.NewEfiPay(credentials)

	body := map[string]interface{} {}	
	const chave = "" //sua chave Pix Efí
	
	

	res, err := efi.PixDeleteWebhook(chave, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
