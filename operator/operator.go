package operator

import (
	"fmt"
	"os/exec"
)

type Interface struct {
	Address    string
	Name       string
	ListenPort int
}

type Operator struct{}

func (o *Operator) ShowInterface() *Interface {
	// TODO: get wg interface via wg cmd
	// TODO: get address via ip cmd
	cmd := exec.Command("wg")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("コマンド実行中にエラーが発生しました:", err)
		return nil
	}
	fmt.Println(string(out))
	return &Interface{}
}

// TODO: implement
func (o *Operator) GetUsers() {}
