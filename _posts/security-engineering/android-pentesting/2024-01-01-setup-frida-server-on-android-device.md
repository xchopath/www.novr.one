---
layout: post
title: "Setup Frida on Android Device (Frida-Server)"
date: 2024-01-01 08:00:00 +0700
categories: "Android-Pentesting"
---

Tool yang akan digunakan:
- `frida` - Install pada PC/Laptop, digunakan untuk _connect_ ke frida-server.
- `frida-server` - Install pada sistem Android (Root), berfungsi untuk mengkoneksikan Android dan PC/Laptop.

## 1. Frida

Install Frida Tool pada PC/Laptop dengan versi spesifik:
```
virtualenv frida
source frida/bin/activate
pip3 install frida-tools frida==16.1.10
```

## 2. Frida Server

### Download Frida Server
<https://github.com/frida/frida/releases/tag/16.1.10>

Catatan:
- Versi harus sesuai dengan versi Frida / Frida Tool yang digunakan.
- Arsitektur harus sesuai dengan arsitektur Andorid yang digunakan (misal `android-arm64`).

### Copy to Android Device via ADB

**Note:** Perangkat harus dalam kondisi **Rooted**.

Setup:
```
unxz frida-server-16.1.10-android-arm64.xz
adb push frida-server-16.1.10-android-arm64 /data/local/tmp/
echo -n "chmod +x /data/local/tmp/frida-server-16.1.10-android-arm64" | adb shell
```

Run:
```
adb shell
/data/local/tmp/frida-server-16.1.10-android-arm64
```

## 3. Connection Test

Melihat device yang sudah terhubung dengan Frida (PC/Laptop).

```
frida-ls-devices
```
