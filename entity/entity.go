package entity

type NetworkInterface struct {
	Name       string
	Address    string
	ListenPort int
}

type User struct {
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Keys  UserKeys `json:"keys"`
}

type UserKeys struct {
	PublicKey    string `json:"public_key"`
	PrivateKey   string `json:"private_key"`
	PresharedKey string `json:"preshared_key"`
}
