# Wildcard Recon 101

**Installation**

Subfinder:
```sh
curl -s "https://api.github.com/repos/projectdiscovery/subfinder/releases/latest" | grep "subfinder_.*_linux_amd64.zip" | cut -d : -f 2,3 | tr -d \" | wget -qi -
unzip -o $(ls | grep subfinder_ | grep zip$)
chmod +x subfinder
mv subfinder /usr/local/bin/
rm $(ls | grep subfinder_ | grep zip$)
```

DNSX:
```sh
curl -s "https://api.github.com/repos/projectdiscovery/dnsx/releases/latest" | grep "dnsx_.*_linux_amd64.zip" | cut -d : -f 2,3 | tr -d \" | wget -qi -
unzip -o $(ls | grep dnsx_ | grep zip$)
chmod +x dnsx
mv dnsx /usr/local/bin/
rm $(ls | grep dnsx_ | grep zip$)
```

```sh
curl -s "https://api.github.com/repos/projectdiscovery/httpx/releases/latest" | grep "httpx_.*_linux_amd64.zip" | cut -d : -f 2,3 | tr -d \" | wget -qi -
unzip -o $(ls | grep httpx_ | grep zip$)
chmod +x httpx
mv httpx /usr/local/bin/
rm $(ls | grep httpx_ | grep zip$)
```

## 1. Subdomain Enumeration

### Instant

Favorite tools:
- Subfinder
- DNSX

**Usage**
```sh
subfinder -all -d domain.com -silent | dnsx -resp -a -aaaa -cname -cdn -silent
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/7be2f3f6-8492-4777-a7dc-54b375351ef8)


## 2. HTTP Discovery

Favorite tools:
- HTTPX

**Usage**

```sh
cat domain_list.txt | httpx -title -ports 80,81,443,5000,5601,7500,8000,8080,8090,8081,8443,9000,9080,9090,9100,9200,9300,10000,10443 -status-code -content-length -content-type -ip -favicon -store-response-dir ./httpx-response -silent
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/20ea734c-19b5-4bfd-80d4-62eb3cb6121d)

## 3. Nuclei HTTP Vulnerability Scan

```sh
nuclei -t http/exposures,http/cves,http/vulnerabilities,http/misconfiguration,http/default-logins,http/exposed-panels,http/miscellaneous,http/technologies -silent
```
