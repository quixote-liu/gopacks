package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := newConfig("a", 1)
	fmt.Println("before:", a)
	fmt.Println("address:", a.NetworkConfig[0].Addresses[0].IPNetmask)
	fmt.Println("len:", len(a.NetworkConfig))
	b := Config{
		PxeMtu:     2,
		ServiceMtu: 2,
		NetworkConfig: []NetworkConfig{
			{Type: "b", Addresses: []Addresses{newAddresses("b")}},
			{Type: "bb"},
		},
	}
	bb, _ := json.Marshal(b)
	fmt.Printf("=================\n bb.String()=%s\n=================\n", string(bb))
	_ = json.Unmarshal(bb, &a)
	fmt.Println("after:", a)
	fmt.Println("address:", a.NetworkConfig[0].Addresses[0].IPNetmask)
	fmt.Println("len:", len(a.NetworkConfig))
}

func newConfig(u string, n int) Config {
	return Config{
		PxeMtu:        n,
		ServiceMtu:    n,
		NetworkConfig: []NetworkConfig{newNetworkConfig(u)},
	}
}

type Config struct {
	PxeMtu        int             `yaml:"pxe_mtu,omitempty" json:"pxe_mtu,omitempty"`
	ServiceMtu    int             `yaml:"service_mtu,omitempty" json:"service_mtu,omitempty"`
	NetworkConfig []NetworkConfig `yaml:"network_config,omitempty" json:"network_config,omitempty"`
}

func newNetworkConfig(u string) NetworkConfig {
	return NetworkConfig{
		Type:      u,
		Mtu:       u,
		Name:      u,
		Addresses: []Addresses{newAddresses(u)},
		Members:   []Members{newMember(u)},
		Routes:    []Routes{newRoutes(u)},
	}
}

type NetworkConfig struct {
	Type      string      `yaml:"type,omitempty" json:"type,omitempty"`
	Mtu       string      `yaml:"mtu,omitempty" json:"mtu,omitempty"`
	Name      string      `yaml:"name,omitempty" json:"name,omitempty"`
	Addresses []Addresses `yaml:"addresses,omitempty" json:"addresses,omitempty"`
	Members   []Members   `yaml:"members,omitempty" json:"members,omitempty"`
	Routes    []Routes    `yaml:"routes,omitempty" json:"routes,omitempty"`
}

func newAddresses(u string) Addresses {
	return Addresses{IPNetmask: u}
}

type Addresses struct {
	IPNetmask string `yaml:"ip_netmask,omitempty" json:"ip_netmask,omitempty"`
}

func newMember(u string) Members {
	return Members{
		Type:           u,
		Name:           u,
		BondingOptions: u,
		Mtu:            u,
		Members:        []Members{{Type: u, Name: u}},
		Addresses:      []Addresses{newAddresses(u)},
	}
}

type Members struct {
	Type           string      `yaml:"type,omitempty" json:"type,omitempty"`
	Name           string      `yaml:"name,omitempty" json:"name,omitempty"`
	BondingOptions string      `yaml:"bonding_options,omitempty" json:"bonding_options,omitempty"`
	Mtu            string      `yaml:"mtu,omitempty" json:"mtu,omitempty"`
	Members        []Members   `yaml:"members,omitempty" json:"members,omitempty"`
	VlanID         int         `yaml:"vlan_id,omitempty" json:"vlan_id,omitempty"`
	Addresses      []Addresses `yaml:"addresses,omitempty" json:"addresses,omitempty"`
}

func newRoutes(u string) Routes {
	return Routes{
		NextHop: "a",
	}
}

type Routes struct {
	NextHop string `yaml:"next_hop,omitempty" json:"next_hop,omitempty"`
	Default bool   `yaml:"default,omitempty" json:"default,omitempty"`
}
