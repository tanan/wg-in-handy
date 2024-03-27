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
	cmd := exec.Command("wg", "show", "wg0", "listen-port")
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

func (o *Operator) createKey() (entity.AuthKeys, error) {
	// wg genkey | tee privatekey | wg pubkey > publickey
	genPrivKeyCmd := exec.Command("wg", "genkey")
	privateKey, _ := genPrivKeyCmd.CombinedOutput()
	pubKeyCmd := exec.Command("wg", "pubkey")
	pubKeyCmd.Stdin = strings.NewReader(string(privateKey))
	publicKey, _ := pubKeyCmd.CombinedOutput()
	preSharedKeyCmd := exec.Command("wg", "genpsk")
	preSharedKey, _ := preSharedKeyCmd.CombinedOutput()
	return entity.AuthKeys{
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

	user.AuthKeys, _ = o.createKey()

	b, _ := json.MarshalIndent(user, "", "  ")
	_, err = f.Write(b)
	if err != nil {
		slog.Error("can't write content", slog.String("file", fileName))
		return err
	}
	return nil
}

func (o *Operator) GenerateServerConfig(networkInterface entity.NetworkInterface, routes []entity.Route, users []entity.User) error {
	f, _ := os.Create("wg0.conf.sample")
	defer f.Close()

	var row []string
	row = append(row, "[Interface]")
	row = append(row, "PrivateKey = "+networkInterface.AuthKeys.PrivateKey)
	row = append(row, "Address = "+networkInterface.Address)
	row = append(row, "ListenPort = "+strconv.Itoa(networkInterface.ListenPort))
	f.Write([]byte(strings.Join(row, "\n")))

	f.Write([]byte("\n\n"))

	for _, v := range users {
		var row []string
		row = append(row, "[Peer]")
		row = append(row, "PublicKey = "+v.AuthKeys.PublicKey)
		row = append(row, "AllowedIPs = "+o.toStringFromRoutes(routes))
		f.Write([]byte(strings.Join(row, "\n")))
	}

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
