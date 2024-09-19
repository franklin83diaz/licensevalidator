package main

import (
	"flag"
	"fmt"
	"licensevalidator"
	"log"
)

func main() {

	licenseServer := flag.String("server", "127.0.0.1", "License server address")
	serialNumber := flag.String("serial", "", "Serial number")

	flag.Parse()

	protectedId, err := licensevalidator.GetId("myapp")
	if err != nil {
		log.Fatal(err)
	}
	//Just to show the protected ID, normally just used internally
	fmt.Println("Protected ID:", protectedId)

	ok, err := licensevalidator.Validate(*licenseServer, *serialNumber, "myapp", "app.lic", "pubkey.pem", "privkey.pem")
	if err != nil {
		log.Fatal(err)
	}

	if ok {
		fmt.Println("License is valid")
	} else {
		fmt.Println("License is not valid")
	}

}
