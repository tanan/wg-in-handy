package operator

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/tanan/wg-in-handy/entity"
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
	cmd := exec.Command("sudo", "wg", "show", "wg0", "listen-port")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.Trim(string(string(out)), "\n"))
}

// TODO: Create server conf file (include user list)

// TODO: implement
func (o *Operator) GetUsers() {

}

func (o *Operator) CreateUser(user *entity.User) error {
	fileName := fmt.Sprintf("%s/%s.json", "/etc/wireguard/client", user.Name)
	f, err := os.Create(fileName)
	if err != nil {
		slog.Error("can't create a file", slog.String("file", fileName))
		return err
	}
	defer f.Close()
	b, _ := json.Marshal(user)
	_, err = f.Write(b)
	if err != nil {
		slog.Error("can't write content", slog.String("file", fileName))
		return err
	}
	return nil
}
