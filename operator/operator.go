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

func (o *Operator) createKey() (entity.UserKeys, error) {
	// wg genkey | tee privatekey | wg pubkey > publickey
	genPrivKeyCmd := exec.Command("sudo", "wg", "genkey")
	privateKey, _ := genPrivKeyCmd.CombinedOutput()
	pubKeyCmd := exec.Command("sudo", "wg", "pubkey")
	pubKeyCmd.Stdin = strings.NewReader(string(privateKey))
	publicKey, _ := pubKeyCmd.CombinedOutput()
	preSharedKeyCmd := exec.Command("sudo", "wg", "genpsk")
	preSharedKey, _ := preSharedKeyCmd.CombinedOutput()
	return entity.UserKeys{
		PublicKey:    strings.Trim(string((publicKey)), "\n"),
		PrivateKey:   strings.Trim(string(privateKey), "\n"),
		PresharedKey: strings.Trim(string(preSharedKey), "\n"),
	}, nil
}

func (o *Operator) CreateUser(user *entity.User) error {
	fileName := fmt.Sprintf("%s/%s.json", "/etc/wireguard/client", user.Name)
	f, err := os.Create(fileName)
	if err != nil {
		slog.Error("can't create a file", slog.String("file", fileName))
		return err
	}
	defer f.Close()

	user.Keys, _ = o.createKey()

	b, _ := json.MarshalIndent(user, "", "  ")
	_, err = f.Write(b)
	if err != nil {
		slog.Error("can't write content", slog.String("file", fileName))
		return err
	}
	return nil
}
