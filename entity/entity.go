package entity

import (
	"net"
	"net/netip"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Address  string   `json:"address"`
	AuthKeys AuthKeys `json:"auth_keys"`
}

type AuthKeys struct {
	PublicKey    string `json:"public_key"`
	PrivateKey   string `json:"private_key"`
	PresharedKey string `json:"preshared_key"`
}

type Route struct {
	Address     string
	Description string
}

type NetworkInterface struct {
	Name       string
	Address    netip.Addr
	Network    net.IPNet
	Endpoint   string
	ListenPort int
	AuthKeys   AuthKeys
}

func NewNetworkInterface(name, network, endpoint string, listenPort int, authKeys AuthKeys) (*NetworkInterface, error) {
	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		return nil, err
	}
	addr := netip.MustParseAddr(ip.To4().String()).Next()
	return &NetworkInterface{
		Name:       name,
		Address:    addr,
		Network:    *ipnet,
		Endpoint:   endpoint,
		ListenPort: listenPort,
		AuthKeys:   authKeys,
	}, nil
}
