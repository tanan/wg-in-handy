package entity

type PublicKey string

type NetworkInterface struct {
	Name       string
	Address    string
	ListenPort int
}

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PublicKey PublicKey `json:"public_key"`
}
