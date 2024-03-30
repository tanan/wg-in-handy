package operator

import (
	"log/slog"
	"net"
	"net/netip"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tanan/wg-in-handy/entity"
)

func (o *Operator) ShowWGInterface() *entity.NetworkInterface {
	// TODO: get wg interface via wg cmd
	addr, network, err := o.getAddress(WireGuardInterface)
	if err != nil {
		slog.Error("Failed to get interface address", slog.String("error", err.Error()))
	}
	port, err := o.getListenPort()
	if err != nil {
		slog.Error("Failed to get listen-port", slog.String("error", err.Error()))
	}
	return &entity.NetworkInterface{
		Name:       WireGuardInterface,
		Address:    addr,
		Network:    network,
		ListenPort: port,
	}
}

func (o *Operator) getAddress(inf string) (netip.Addr, net.IPNet, error) {
	cmd := exec.Command("ip", "-f", "inet", "-o", "addr", "show", inf)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return netip.Addr{}, net.IPNet{}, err
	}
	res := strings.Split(string(out), " ")
	ip, ipnet, err := net.ParseCIDR(res[AddressNum])
	if err != nil {
		return netip.Addr{}, net.IPNet{}, err
	}
	return netip.MustParseAddr(ip.To4().String()), *ipnet, nil
}

func (o *Operator) getListenPort() (int, error) {
	cmd := exec.Command("wg", "show", WireGuardInterface, "listen-port")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Trim(string(string(out)), "\n"))
}

func (o *Operator) UpWGInterface() error {
	filename := "/etc/wireguard/wg0.conf"
	if _, err := os.Stat(filename); err != nil {
		slog.Error("file does not exist", slog.String("filename", filename))
		return err
	}

	cmd := exec.Command("wg-quick", "up", WireGuardInterface)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}

func (o *Operator) DownWGInterface() error {
	filename := "/etc/wireguard/wg0.conf"
	if _, err := os.Stat(filename); err != nil {
		slog.Error("file does not exist", slog.String("filename", filename))
		return err
	}

	cmd := exec.Command("wg-quick", "down", WireGuardInterface)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
