package operator

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/tanan/wg-in-handy/entity"
)

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
