# Konflik Routing antara VPN dan Local Network

### Troubleshoot

Problem:

```sh
traceroute <DEST IP>
```
```
traceroute to <DEST IP> (<DEST IP>), 30 hops max, 60 byte packets
 1  192.168.100.1 (192.168.100.1)  4.696 ms  4.542 ms  4.742 ms
 2  10.154.128.1 (10.154.128.1)  5.081 ms  6.259 ms  6.338 ms
 3  180.252.1.249 (180.252.1.249)  6.626 ms  6.815 ms  8.564 ms
 4  192.168.144.2 (192.168.144.2)  8.490 ms  8.415 ms  8.342 ms
 5  192.168.144.1 (192.168.144.1)  8.250 ms  8.177 ms  8.078 ms
 6  * * *
 7  * * *
 8  * * *
 9  * * *
10  * * *
```


Check Route:

```sh
ip route
```

Binding Interface:

```sh
sudo traceroute -i <VPN INTERFACE> <DEST IP>
curl -v --interface <VPN INTERFACE> https://<DEST IP>
```
```
traceroute to <DEST IP> (<DEST IP>), 30 hops max, 60 byte packets
 1  * * *
 2  192.168.5.222 (192.168.5.222)  7.231 ms  7.778 ms  7.801 ms
 3  172.18.20.67 (<DEST IP>)  8.028 ms  8.012 ms  10.919 ms
```

Fix
```sh
sudo ip route add <DEST IP>/32 dev <VPN INTERFACE>
```

Test again:
```sh
traceroute 172.18.20.67
```
```
traceroute to 172.18.20.67 (172.18.20.67), 30 hops max, 60 byte packets
 1  * * *
 2  192.168.5.222 (192.168.5.222)  8.628 ms  8.881 ms  8.909 ms
 3  172.18.20.67 (172.18.20.67)  9.018 ms  9.123 ms  13.677 ms
```
