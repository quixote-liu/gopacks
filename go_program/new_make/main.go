package main

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
)

func main() {
	client := new(gophercloud.ProviderClient)

	fmt.Println("client.EndpointLocator == nil", client.EndpointLocator == nil)
}
