package operator

import (
	"os/exec"
	"strings"

	"github.com/tanan/wg-in-handy/entity"
)

const (
	NameNum            = 1
	AddressNum         = 6
	WireGuardInterface = "wg0"
)

type Operator struct{}

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
