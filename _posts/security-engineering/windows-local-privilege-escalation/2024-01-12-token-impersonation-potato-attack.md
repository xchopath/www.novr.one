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

## Why Impersonation Exists?

Impersonation adalah kemampuan untuk menjalankan sebuah tugas menggunakan hak istimewa atau identitas pengguna lain. Kita ambil contoh Service IIS (Web Server), ketika User yang menjalankan service IIS bukanlah `NT AUTHORITY\SYSTEM` (LocalSystem), proses tersebut mungkin tidak memiliki privilege yang cukup untuk mengakses beberapa sumber daya di sistem. Namun, dengan menggunakan impersonasi, layanan masih dapat mencoba menjalankan tugas tertentu dengan hak istimewa yang lebih tinggi dan layanan masih dapat melanjutkan tugasnya tanpa harus berhenti hanya karena pengguna bukan LocalSystem.

Jadi yang perlu digaris bawahi yaitu User masih dapat mengakses sumber daya tertentu meskipun tanpa Role yang setara.

<br/>

# Impersonate Privileges

Mari kita bahas dengan lebih sederhana, kita akan membahas terkait privilege `SeImpersonatePrivilege` dan `SeAssignPrimaryTokenPrivilege`.

SeImpersonatePrivilege adalah hak istimewa di Windows yang memungkinkan kita untuk menyamar sebagai pengguna lain (Impersonation) untuk menjalankan sebuah tugas di bawah naungan user tersebut, begitu juga dengan SeAssignPrimaryTokenPrivilege. Jika kita memiliki SeImpersonatePrivilege, kita dapat menggunakan fungsi `CreateProcessWithTokenW()` untuk membuat proses baru dengan token yang kita miliki. Sebagai alternatifnya, SeAssignPrimaryTokenPrivilege memungkinkan kita untuk menggunakan fungsi `CreateProcessAsUserA()`, yang memiliki fungsi serupa.

<br/>

# Potato Attacks

## 1. RottenPotato - Abuse Impersonation through `CoGetInstanceFromIStorage() function`.

<https://learn.microsoft.com/en-us/windows/win32/api/objbase/nf-objbase-cogetinstancefromistorage>
```
| Windows
--| Apps
----| Win32
------| API
--------| Component Object Model (COM)
----------| Objbase.h
------------| CoGetInstanceFromIStorage() function
```
