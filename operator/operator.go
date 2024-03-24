package operator

import (
	"log/slog"
	"os/exec"
	"strconv"
	"strings"
)

type Interface struct {
	Address    string
	Name       string
	ListenPort int
}

const (
	NameNum       = 1
	AddressNum    = 6
	InterfaceName = "wg0"
)

type Operator struct{}

func (o *Operator) ShowInterface() *Interface {
	// TODO: get wg interface via wg cmd
	addr, err := o.getAddress(InterfaceName)
	if err != nil {
		slog.Error("Failed to get interface address", slog.String("error", err.Error()))
	}
	port, err := o.getListenPort()
	if err != nil {
		slog.Error("Failed to get listen-port", slog.String("error", err.Error()))
	}
	return &Interface{
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
	cmd := exec.Command("sudo", "wg", "show", "wg0", "listen-port")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Trim(string(string(out)), "\n"))
}

// TODO: implement
func (o *Operator) GetUsers() {}
