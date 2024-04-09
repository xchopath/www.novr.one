# Wildcard Recon 101

### Installation

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

HTTPX:
```sh
curl -s "https://api.github.com/repos/projectdiscovery/httpx/releases/latest" | grep "httpx_.*_linux_amd64.zip" | cut -d : -f 2,3 | tr -d \" | wget -qi -
unzip -o $(ls | grep httpx_ | grep zip$)
chmod +x httpx
mv httpx /usr/local/bin/
rm $(ls | grep httpx_ | grep zip$)
```

Nuclei:
```sh
curl -s "https://api.github.com/repos/projectdiscovery/nuclei/releases/latest" | grep "nuclei_.*_linux_amd64.zip" | cut -d : -f 2,3 | tr -d \" | wget -qi -
unzip -o $(ls | grep nuclei_ | grep zip$)
chmod +x nuclei
mv nuclei /usr/local/bin/
rm $(ls | grep nuclei_ | grep zip$)
```

FFuF:
```sh
curl -s "https://api.github.com/repos/ffuf/ffuf/releases/latest" | grep "ffuf_.*_linux_amd64.tar.gz" | cut -d : -f 2,3 | tr -d \" | wget -qi -
tar -xzvf $(ls | grep ffuf_ | grep 'tar\.gz'$)
chmod +x ffuf
mv ffuf /usr/local/bin
rm $(ls | grep ffuf_ | grep 'tar\.gz'$)
```

<br/>

## 1. Subdomain Enumeration

### Instant

Favorite tools:
- Subfinder
- DNSX

**Usage**
```sh
subfinder -all -d domain.com -silent | dnsx -resp -a -aaaa -cname -cdn -asn -silent
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/7be2f3f6-8492-4777-a7dc-54b375351ef8)

<br/>

## 2. HTTP Discovery

Favorite tools:
- HTTPX

**Usage**

```sh
cat domain_list.txt | httpx -title -ports 80,81,443,5000,5601,7500,8000,8080,8090,8081,8443,9000,9080,9090,9100,9200,9300,10000,10443 -status-code -content-length -content-type -ip -favicon -store-response-dir ./httpx-response -silent
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/20ea734c-19b5-4bfd-80d4-62eb3cb6121d)

<br/>

## 3. Nuclei HTTP Vulnerability Scan

Tools:
- Nuclei

**Usage**

```sh
nuclei -t http/exposures,http/cves,http/vulnerabilities,http/misconfiguration,http/default-logins,http/exposed-panels,http/miscellaneous,http/technologies -silent
```

<br/>

## 4. Directory Fuzzing

**Craft Custom Wordlist**

```sh
curl -s "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/raft-medium-words.txt" >> /tmp/fuzzlist-raft-medium-words.txt
curl -s "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/raft-medium-files.txt" >> /tmp/fuzzlist-raft-medium-files.txt
curl -s "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/raft-medium-directories.txt" >> /tmp/fuzzlist-raft-medium-directories.txt
curl -s "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Discovery/Web-Content/api/api-endpoints.txt" >> /tmp/fuzzlist-api-endpoints.txt
curl -s "https://raw.githubusercontent.com/maurosoria/dirsearch/master/db/dicc.txt" >> /tmp/fuzzlist-dirsearch.txt
cat /tmp/fuzzlist-*.txt | grep -v '%EXT%' | sort -u >> custom-dirlist.txt
rm /tmp/fuzzlist-*.txt
sudo mv custom-dirlist.txt /usr/share/
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/a485a590-be41-4fb5-929f-961a13fe5728)

**Usage**

Basic

```sh
ffuf -u https://target.com/FUZZ -w /usr/share/custom-dirlist.txt -t 100
```
