package operator

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/tanan/wg-in-handy/entity"
)

func (o *Operator) GenerateServerConfig(networkInterface entity.NetworkInterface, routes []entity.Route, users []entity.User) error {
	f, _ := os.Create("wg0.conf.sample")
	defer f.Close()

	var row []string
	row = append(row, "[Interface]")
	row = append(row, "PrivateKey = "+networkInterface.AuthKeys.PrivateKey)
	row = append(row, "Address = "+networkInterface.Address.String())
	row = append(row, "ListenPort = "+strconv.Itoa(networkInterface.ListenPort))
	f.Write([]byte(strings.Join(row, "\n")))

	f.Write([]byte("\n\n"))

	for _, v := range users {
		var row []string
		row = append(row, "[Peer]")
		row = append(row, "PublicKey = "+v.AuthKeys.PublicKey)
		row = append(row, "PresharedKey = "+networkInterface.AuthKeys.PresharedKey)
		row = append(row, "AllowedIPs = "+o.toStringFromRoutes(routes))
		f.Write([]byte(strings.Join(row, "\n")))
	}

	f.Write([]byte("\n"))

	return nil
}

func (o *Operator) GenerateClientConfig(networkInterface entity.NetworkInterface, routes []entity.Route, user entity.User) error {
	f, _ := os.Create(fmt.Sprintf("%s.conf.sample", user.Name))
	defer f.Close()

	var row []string
	row = append(row, "[Interface]")
	row = append(row, "PrivateKey = "+user.AuthKeys.PrivateKey)
	row = append(row, "Address = "+user.Address)
	row = append(row, "\n")
	row = append(row, "[Peer]")
	row = append(row, "PublicKey = "+networkInterface.AuthKeys.PrivateKey)
	row = append(row, "PresharedKey = "+networkInterface.AuthKeys.PresharedKey)
	row = append(row, "EndPoint = "+networkInterface.Endpoint)
	row = append(row, "AllowedIPs = "+o.toStringFromRoutes(routes))
	f.Write([]byte(strings.Join(row, "\n")))

	f.Write([]byte("\n"))

	return nil
}

func (o *Operator) toStringFromRoutes(routes []entity.Route) string {
	var row []string
	for _, v := range routes {
		row = append(row, v.Address)
	}
	return strings.Join(row, ",")
}
