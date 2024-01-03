---
layout: post
title: "Bypass Root Detection with Hook Method using Frida (AndroGoat.apk Study-case)"
date: 2024-01-02 10:00:00 +0700
categories: "Android-Pentesting"
---

![Frida Hook](https://github.com/xchopath/www.novr.one/assets/44427665/612d3adb-bdb0-4f89-90ba-534f0c0739d2)


Tools need to be prepared:

|     Tool     | Install on | Purpose                                                |
|:-------------|:----------:|--------------------------------------------------------|
|   jadx-gui   | PC/Laptop  | Reverse Engineering file APK                           |
|     frida    | PC/Laptop  | Untuk menginjeksi script untuk melakukan Hook          |
| frida-server |  Android   | Membuka koneksi Android agar Frida PC dapat terkoneksi |

<br/>

## Lab Setup (AndroGoat App)

Download:
- <https://github.com/satishpatnayak/MyTest/blob/master/AndroGoat.apk>

<br/>

-----

<br/>

# 1. APK Code Analysis with JADX

Jalankan `jadx-gui` dan buka file APK melalui jadx untuk melakukan _decompile_.

## Find a Function

![jadx-gui androgoat](https://github.com/xchopath/www.novr.one/assets/44427665/1f762cda-a33b-4e2c-a95e-e6e895ada35d)

Kueri-kueri yang dapat digunakan untuk menemukan fungsi _Root Detection_.
- `bin/su`
- `"su"`
- `/system/`
- `/data/data`
- `rootdetect`
- `isroot`

![jadx search](https://github.com/xchopath/www.novr.one/assets/44427665/065af768-2d93-4d65-bd5d-a951ee098956)

<br/>

## Code Analysis

Pada tahap ini, kita perlu melakukan pembedahan terkait fungsi-fungsi apa saja yang digunakan.

Ditemukan fungsi `isRooted()`

![Fungsi isRoot](https://github.com/xchopath/www.novr.one/assets/44427665/4889768f-27c0-4c44-aacd-b3e3c9514af8)

Ditemukan lagi fungsi `isRooted1()`

![image](https://github.com/xchopath/www.novr.one/assets/44427665/f85fb369-7187-4e2e-a753-f72382e82368)

Di sini kita telah mengetahui bahwa terdapat 2 fungsi yang digunakan pada modul pendeteksi Root-nya.

Karena `return` yang diharapkan itu `false`, maka perlu dicatat juga akan hal ini.

<br/>

## Hook Logic

Berikut ini adalah logika yang akan kita digunakan untuk _Hooking_.
```
isRooted()  => True  => Root Detected
isRooted1() => True  => Root Detected
isRooted()  => False => Not Rooted
isRooted1() => False => Not Rooted
```

<br/>

**Package Location**

Perlu diingat juga bahwa Code yang kita analisa tadi itu terdapat di dalam _package_ `owasp.sat.agoat.RootDetectionActivity`, hal ini yang akan digunakan pada Script Hook nantinya.

![owasp sat agoat RootDetectionActivity](https://github.com/xchopath/www.novr.one/assets/44427665/d0a5e098-cd06-4448-a794-24942aa77ebb)

<br/>

-----

<br/>

# 2. Hooking

> Hooking:
> Memanipulasi aktivitas (secara _Real-Time_) pada sebuah aplikasi Android yang sedang berjalan.

**Before Hook**

![Before Hook](https://github.com/xchopath/www.novr.one/assets/44427665/590c510e-2fc3-4ea7-8eb4-d24046fec6e1)

Sebelum dilanjutkan, kita perlu [Setup Frida](https://www.novr.one/security-engineering/android-pentesting/2024-01-01-setup-frida-server-on-android-device) terlebih dahulu.

<br/>

## Run the App and Get App Identifier

```sh
frida-ps -Uai
```

![Get App Identifier](https://github.com/xchopath/www.novr.one/assets/44427665/4ad10cec-1e36-4aa8-bd1c-8327a5808787)

<br/>

## Prepare Script to Hook

Script for example `androgoat-root-bypass.js`:
```javascript
Java.perform(
  function () {
    let CallJavaPackage = Java.use("owasp.sat.agoat.RootDetectionActivity");
    CallJavaPackage["isRooted"].implementation = function() { return false; };
    CallJavaPackage["isRooted1"].implementation = function() { return false; };
  }
);
```

<br/>

## Hook the Activity

```sh
frida -l <hook-script> -U -f <app-identifier>
```

![frida hook](https://github.com/xchopath/www.novr.one/assets/44427665/42185d51-517d-4ed3-9532-7bba73e799a8)

**When the frida script running**

![not rooted](https://github.com/xchopath/www.novr.one/assets/44427665/625dcaf7-3c74-44e6-a733-b98daf010332)
