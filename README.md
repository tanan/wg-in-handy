# wg-in-handy

- create public and private key

```
wg genkey | tee privatekey | wg pubkey > publickey
```

- set configuration

```
$ wg setconf wg0 wg0.conf
```

- get address via ip command

```
ip -f inet -o addr show wg0|cut -d\  -f 7
```
