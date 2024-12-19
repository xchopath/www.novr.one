---
title: "[OSCP] Initial Access - Checklist That Guarantees You to Pwn 'Em All"
author: "novran"
date: 2024-11-27 00:00:00 +0700
categories: ["OSCP Challenge"]
tags: ["Windows Pentest", "Linux Pentest", "FTP Pentest", "SMB Pentest", "Web Pentest", "SNMP Pentest", "OSCP Challenge"]
mermaid: true
image:
  path: /images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all.png
---

## Enumeration

### Port Scan

- [RustScan](https://github.com/RustScan/RustScan)

```bash
rustscan --scripts none -a <TARGET>
```

- [UDPX](https://github.com/nullt3r/udpx)

```bash
udpx -c 500 -w 1000 -t <TARGET>
```

### Web Enumeration

- [Feroxbuster](https://github.com/epi052/feroxbuster) and [SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/directory-list-2.3-medium.txt)

```bash
feroxbuster -C 404 --auto-tune -k --wordlist /usr/share/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt --threads 100 --depth 2 -u <TARGET>
```

- [Dirsearch](https://github.com/maurosoria/dirsearch)

```bash
PYTHONWARNINGS="ignore" python3 ~/dirsearch/dirsearch.py -u <TARGET>
```

# Initial Access

## FTP (21/tcp)

### FTP Anonymous Login

- [NetExec](https://github.com/Pennyw0rth/NetExec)

```bash
netexec ftp <TARGET> -u 'anonymous' -p 'anonymous'
```

### FTP Default Credential Brute Force

[Hydra](https://www.kali.org/tools/hydra/) anf [SecLists/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt)

```bash
hydra -C /usr/share/SecLists/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt <TARGET> ftp
```

### Other FTP Cheatsheet

Download file recursively.

```bash
wget -m --user='<USERNAME>' --password='<PASSWORD>' ftp://<TARGET> --no-passive-ftp
```

## SMB (445/tcp)

Ada dua metode initial access yang dapat dicoba pada layanan SMB, yaitu "Null Session" dan "Guest Login".

Untuk memeriksa service SMB, di sini saya menyarankan untuk menggunakan [NetExec](https://github.com/Pennyw0rth/NetExec).

### SMB Null Session

```bash
netexec smb <TARGET> -u '' -p ''
```

### SMB Guest Login

```bash
netexec smb <TARGET> -u 'guest' -p ''
```

### Other SMB Cheatsheet

Download semua file di dalam SMB secara recursive.

```bash
nxc smb <TARGET> -u 'user' -p 'pass' -M spider_plus -o DOWNLOAD_FLAG=True
```

## SNMP (161/udp)

Preparation:
```bash
sudo apt-get install snmp-mibs-downloader -y
sudo download-mibs
```

Initial Access pada SNMP biasanya dilakukan untuk memperoleh kredensial yang tercatat di dalamnya. Dalam SNMP, kita menggunakan "community string" sebagai input (biasanya bernama "public").

### Get Info use Public Community String

```bash
snmpbulkwalk -c public -v2c <TARGET> NET-SNMP-EXTEND-MIB::nsExtendOutputFull
```

Namun, jika community string bernama "public" tidak tersedia, kita perlu melakukan Brute Force untuk menemukannya.

### Brute Force Community String

Terdapat dua wordlist umum yang dapat digunakan, yaitu SecLists atau Wordlist dari Metasploit Framework.

```bash
onesixtyone -c /usr/share/SecLists/Discovery/SNMP/snmp.txt <TARGET>
```

```bash
onesixtyone -c /usr/share/metasploit-framework/data/wordlists/snmp_default_pass.txt <TARGET>
```

SNMP Walk menggunakan custom community string.

```bash
snmpbulkwalk -c <COMMUNITY_STRING_NAME> -v2c <TARGET> NET-SNMP-EXTEND-MIB::nsExtendOutputFull
```

## Found Weird Port? (Uncommon Services)

Google it.

```
port XXX exploit
```

## Web Application (HTTP)

### Enumerate directory and files

### Found an underrated CMS, framework, or application?

Google it
```
<APPLICATION_NAME> exploit
<APPLICATION_NAME_AND_VERSION> exploit
```

### Does the application require login?

- Try `admin:admin`
- or Google it: `<APPLICATION_NAME> default credentials`

## Found Numerous Unnecessary Files?

### Use exiftool to gather username list

Basic:

```
exiftool <FILENAME>
```

Bunch of DOCXs/PDFs

```sh
find . -type f | xargs -I {} exiftool {} | grep ^'Author'
```

And spray it across all services or apps using `<username>:<username>` to log in.

## Found Protected File?

You need to crack it.

![anything 2 john](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-file2john.png)
