package entity

type NetworkInterface struct {
	Name       string
	Address    string
	ListenPort int
	AuthKeys   AuthKeys
}

type Route struct {
	Address     string
	Description string
}

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	AuthKeys AuthKeys `json:"auth_keys"`
}

type AuthKeys struct {
	PublicKey    string `json:"public_key"`
	PrivateKey   string `json:"private_key"`
	PresharedKey string `json:"preshared_key"`
}
