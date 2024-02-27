---
layout: post
title: "Kerberoasting Attack (Active Directory)"
date: 2024-01-07 07:00:00 +0700
categories: "Active-Directory-Pentesting"
---

Kerberoasting adalah salah satu serangan yang mengeksploitasi kerentanan dalam sistem otentikasi Kerberos di Active Directory. Dalam serangan ini, penyerang akan melakukan percobaan untuk mendapatkan tiket layanan (TGS) pada sebuah layanan tertentu yang dijalankan oleh akun layanan (Service Account), seperti database server (mssql) dan lainnya. Jika berhasil mendapatkannya, penyerang akan mencoba untuk membongkar password yang terkandung di dalam tiket tersebut untuk mendapatkan akses ke akun layanan (si pemilik tiket).

**Alur Kerberoasting:**

1. Attacker mengekstrak tiket layanan atau biasa disebut TGS (Ticket Granting Services) dari Kerberos yang bertindak sebagai distributor kunci atau Key Distribution Center (KDC).
2. Ekstraksi TGS dapat dilakukan menggunakan User yang sudah compromised. Biasanya User tersebut sudah memiliki akses untuk mengakses SPN (Service Principal Names).
3. Pada dasarnya TGS itu terenkripsi dengan password milik Service Account (yang akan di-crack nantinya) dan Service Account biasanya memiliki akses khusus ke layanan tertentu (misal: mssql).
4. Ketika Attacker sudah berhasil mendapatkan TGS, kemudian Attacker akan melakukan cracking untuk mendapatkan password (plain-text) milik Service Account yang terkandung di dalam TGS-nya. Hal ini dilakukan dengan menggunakan metode seperti Brute Force secara offline.
5. Jika password yang terkandung pada TGS berhasil di-crack, maka Attacker dapat login menggunakan Credential milik Service Account atau akun si pemilik tiket tersebut.

**Authentication Required! Because this is Post-Compromised Attack!**

Serangan Kerberoasting dapat dilakukan ketika Attacker sudah memiliki User yang berhasil di-takeover, karena untuk mengekstrak TGS dan SPN itu perlu menggunakan `username` dan `password`.

<br/>

# Proof of Concept

Required Tools.

```
sudo apt install impacket-scripts -y
```

<br/>

## Extract TGS and Service Principal Names (SPN) Through Compromised User 

Command
```sh
impacket-GetUserSPNs <domain>/<username>:<password> -request -dc-ip <ip domain controller>
```

![Impacket-GetSPNs TGS](https://github.com/xchopath/www.novr.one/assets/44427665/2ae43848-74da-4a58-b7e6-40d15e37d238)


**Catatan:**

Jika mengalami Error `Kerberos SessionError: KRB_AP_ERR_SKEW(Clock skew too great)` maka kita hanya perlu untuk sinkronisasi waktu dan tanggal pada Host yang kita gunakan dengan Server Active Directory.

Get AD Timestamp via LDAP.
```sh
ldapsearch -LLL -x -H ldap://<ip domain controller> -b '' -s base '(objectclass=*)' | grep currentTime
```

Synchronize the datetime.
```sh
sudo timedatectl set-ntp false
sudo timedatectl set-time "<YYYY-MM-DD HH:MM:SS>"
```

<br/>

## Crack the TGS (Kerberoasting!)

Command
```sh
john --format=krb5tgs --wordlist=<wordlist> <tgs file>
```

![Kerberoasting with john 1](https://github.com/xchopath/www.novr.one/assets/44427665/4f57812d-f2f2-4347-a1d7-8953ee106f99)
![Kerberoasting with john 2](https://github.com/xchopath/www.novr.one/assets/44427665/7fac6097-a141-4f53-ac99-afa03dc1798c)
