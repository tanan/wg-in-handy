package operator

import (
	"log/slog"
	"os/exec"
	"strings"
)

type Interface struct {
	Address    string
	Name       string
	ListenPort int
}

const (
	NameNum    = 1
	AddressNum = 6
)

type Operator struct{}

func (o *Operator) ShowInterface() *Interface {
	// TODO: get wg interface via wg cmd
	cmd := exec.Command("ip", "-f", "inet", "-o", "addr", "show", "ens4")
	out, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("Failed to run command", slog.String("error", err.Error()))
		return nil
	}
	res := strings.Split(string(out), " ")
	return &Interface{
		Name:    res[NameNum],
		Address: res[AddressNum],
	}
}

// TODO: implement
func (o *Operator) GetUsers() {}
