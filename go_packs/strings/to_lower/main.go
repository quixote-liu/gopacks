package main

import (
	"fmt"
	"strings"
)

// var deviceOwnerResourceTypeMapping = map[string]int{
// 	"octavia":                                ResourceTypeOctaviaLoadbalancer,
// 	"compute:nova":                           ResourceTypeComputeServer,
// 	"network:floatingip":                     ResourceTypeNetworkFloatingIP,
// 	"network:router_interface":               ResourceTypeNetworkRouter,
// 	"network:router_interface_distributed":   ResourceTypeNetworkRouter,
// 	"network:ha_router_replicated_interface": ResourceTypeNetworkRouter,
// }

// var deviceOwnerNameMapping = map[string]string{
// 	"network:dhcp":                "DHCP",
// 	"network:router_gateway":      "路由器网关",
// 	"network:router_ha_interface": "路由器VRRP",
// }

func main() {
	caseStr := "NEtwORk:floaTIngip"
	lower := strings.ToLower(caseStr)
	fmt.Println("lower:", lower)
}
