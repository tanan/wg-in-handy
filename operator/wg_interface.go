package operator

import (
	"log/slog"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tanan/wg-in-handy/entity"
)

const WireGuardInterface = "wg0"

func (o *Operator) ShowInterface() *entity.NetworkInterface {
	// TODO: get wg interface via wg cmd
	addr, err := o.getAddress(InterfaceName)
	if err != nil {
		slog.Error("Failed to get interface address", slog.String("error", err.Error()))
	}
	port, err := o.getListenPort()
	if err != nil {
		slog.Error("Failed to get listen-port", slog.String("error", err.Error()))
	}
	return &entity.NetworkInterface{
		Name:       InterfaceName,
		Address:    addr,
		ListenPort: port,
	}
}

func (o *Operator) getAddress(inf string) (string, error) {
	cmd := exec.Command("ip", "-f", "inet", "-o", "addr", "show", inf)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	res := strings.Split(string(out), " ")
	return res[AddressNum], nil
}

func (o *Operator) getListenPort() (int, error) {
	cmd := exec.Command("wg", "show", WireGuardInterface, "listen-port")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Trim(string(string(out)), "\n"))
}

func (o *Operator) CreateWGInterface(wgNetwork string) error {
	// ip link add dev wg0 type wireguard
	cmd := exec.Command("ip", "link", "add", "dev", WireGuardInterface, "type", "wireguard")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	// ip address add dev wg0 192.168.2.1/24
	addrCmd := exec.Command("ip", "address", "add", "dev", WireGuardInterface, wgNetwork)
	_, err = addrCmd.CombinedOutput()
	if err != nil {
		return err
	}

	// wg setconf wg0 myconfig.conf
	wgConfPath := "/etc/wireguard" + "wg0.conf"
	wgCmd := exec.Command("wg", "setconf", WireGuardInterface, wgConfPath)
	_, err = wgCmd.CombinedOutput()
	if err != nil {
		return err
	}

	// ip link set up dev wg0
	linkUpCmd := exec.Command("ip", "link", "set", "up", "dev", WireGuardInterface)
	_, err = linkUpCmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
