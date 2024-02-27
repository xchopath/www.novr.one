---
layout: post
title: "Token Impersonation (Potato Attack)"
date: 2024-01-12 07:00:00 +0700
categories: "Windows-Local-Privilege-Escalation"
---

### TLDR; What is Token?

Ketika kita Login ke dalam sistem, baik itu secara lokal, melalui jaringan, bertindak sebagai Service, atau bahkan langsung memanggil fungsi API LogonUser. Saat itu, paket otentikasi akan membuat sesi Login dan kemudian Local Security Authority (LSA) membuat token akses untuk pengguna tersebut.

Token ini merepresentasikan berisi informasi:
- ID sesi Login
- User dan Group SID
- Integrity level
- Privilege yang dipegang oleh pengguna atau grup tempat pengguna berada

## Why "Impersonation" Exists?

Impersonation adalah kemampuan untuk menjalankan sebuah tugas menggunakan hak istimewa atau identitas pengguna lain. Kita ambil contoh Service IIS (Web Server), ketika User yang menjalankan service IIS bukanlah `NT AUTHORITY\SYSTEM` (LocalSystem), maka proses yang dijalankan oleh service tersebut mungkin tidak memiliki privilege yang cukup untuk mengakses beberapa sumber daya di sistem. Namun, dengan menggunakan Impersonation, layanan IIS masih dapat menjalankan tugas tertentu dengan hak istimewa tanpa role khusus dan layanan masih dapat melanjutkan tugasnya tanpa harus terhenti hanya karena user IIS tidak memiliki Role yang setara dengan LocalSystem.

<br/>

# Impersonate Privileges

Mari kita bahas sedikit teknis terkait privilege `SeImpersonatePrivilege` dan `SeAssignPrimaryTokenPrivilege`.

SeImpersonatePrivilege adalah hak istimewa di Windows yang memungkinkan kita untuk menyamar sebagai User lain (Impersonation) untuk menjalankan sebuah tugas di bawah naungan User tersebut, begitu juga dengan SeAssignPrimaryTokenPrivilege. Jika kita memiliki SeImpersonatePrivilege, kita dapat menggunakan fungsi `CreateProcessWithTokenW()` untuk membuat proses baru dengan token yang kita miliki. Sebagai alternatifnya, SeAssignPrimaryTokenPrivilege memungkinkan kita untuk menggunakan fungsi `CreateProcessAsUserA()`, yang memiliki fungsi serupa.

Windows API Functions:
- [CreateProcessWithTokenW](https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-createprocesswithtokenw) (SeImpersonatePrivilege needed) adalah fungsi dalam pemrograman Windows API yang digunakan untuk membuat proses baru menggunakan token keamanan (security token) dari pengguna atau sesi tertentu.
- [CreateProcessAsUserA](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessasusera) (SeAssignPrimaryTokenPrivilege needed) adalah fungsi dalam Windows API yang mirip dengan CreateProcessWithTokenW. Namun, CreateProcessAsUserA digunakan untuk membuat proses baru sebagai pengguna tertentu.

<br/>

Check privilege using CMD:
```powershell
whoami /priv
```

<br/>

# Impersonation Exploit to Privilege Escalation

![Potato Exploit](https://github.com/xchopath/www.novr.one/assets/44427665/ad2b9f32-a642-4328-8e96-b8f4d20fbb21)

Potato / Impersonation Exploits:
- [RottenPotato](https://github.com/foxglovesec/RottenPotato) / [JuicyPotato](https://github.com/ohpe/juicy-potato) - All windows version until 2018.
- [PrintSpoofer](https://github.com/itm4n/PrintSpoofer) - All windows version until 2019.
- [RoguePotato](https://github.com/antonioCoco/RoguePotato) - All windows version until 2020.
- [GodPotato](https://github.com/BeichenDream/GodPotato) - All windows version until 2022.
