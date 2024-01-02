# SSH Tunneling to Socks5 Server (Local)

Create a Socks5 Server in Local
```
ssh -D 8888 -q -C -N user@yourserver
```

Test Connection
```
curl -s --proxy "socks5://localhost:8888" ipinfo.io
```
