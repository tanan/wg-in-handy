package entity

type PublicKey string

type NetworkInterface struct {
	Name       string
	Address    string
	ListenPort int
}

type User struct {
	Name      string
	Email     string
	PublicKey PublicKey
}
