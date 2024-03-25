# wg-in-handy

- create public and private key

```
$ wg genkey | tee privatekey | wg pubkey > publickey
```

- set configuration

```
$ wg setconf wg0 wg0.conf
```

or

```
$ wg set wg0 listen-port 51820 private-key /root/server/privatekey peer {Peer Public Key} allowed-ips 192.168.2.0/24 endpoint {Public IP}:51820
```

- get address via ip command

```
$ ip -f inet -o addr show wg0|cut -d\  -f 7
```

- create interface

```
$ ip link add dev wg0 type wireguard
$ ip address add dev wg0 192.168.2.1/24
$ ip link set wg0 up
```

- create user

```
curl -XPOST localhost:8080/v1/users -d '
{
    "email": "foo@example.com"
}
'
```