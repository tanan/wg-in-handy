# wg-in-handy

- create public and private key

```
$ wg genkey | tee privatekey | wg pubkey > publickey
```

- set configuration

```
$ wg setconf wg0 wg0.conf
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
