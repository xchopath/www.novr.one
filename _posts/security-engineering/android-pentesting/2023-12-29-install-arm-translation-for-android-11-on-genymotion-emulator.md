---
layout: post
title: "Install ARM Translation for Android 11 on Genymotion (Emulator)"
date: 2023-12-29 00:00:00 +0700
categories: "Android-Pentesting"
---

### 1. Download Library ARM for Android 11

```
git clone https://github.com/xchopath/Genymotion_A11_libhoudini
```

```
cd Genymotion_A11_libhoudini/
```

### 2. Copy Library to Android

Mengatur izin agar direktori root (/) di dalam Android dapat di-write (pastikan ADB sudah terkoneksi dengan device).

```
echo -n "mount -o rw,remount /" | adb shell
```

Copy library system yang baru saja di-download ke dalam Android.

```
adb push system /
```

### 3. Add Configuration

Masuk ke dalam Android (Shell).

```
adb shell
```

Tambahkan konfigurasi untuk mengatur parameter arsitektur CPU dan eksekusi kode natif pada Android System.

```
echo -ne "\nro.product.cpu.abilist=x86_64,x86,arm64-v8a,armeabi-v7a,armeabi\nro.product.cpu.abilist32=x86,armeabi-v7a,armeabi\nro.product.cpu.abilist64=x86_64,arm64-v8a\nro.vendor.product.cpu.abilist=x86_64,x86,arm64-v8a,armeabi-v7a,armeabi\nro.vendor.product.cpu.abilist32=x86,armeabi-v7a,armeabi\nro.vendor.product.cpu.abilist64=x86_64,arm64-v8a\nro.odm.product.cpu.abilist=x86_64,x86,arm64-v8a,armeabi-v7a,armeabi\nro.odm.product.cpu.abilist32=x86,armeabi-v7a,armeabi\nro.odm.product.cpu.abilist64=x86_64,arm64-v8a\nro.dalvik.vm.native.bridge=libhoudini.so\nro.enable.native.bridge.exec=1\nro.enable.native.bridge.exec64=1\nro.dalvik.vm.isa.arm=x86\nro.dalvik.vm.isa.arm64=x86_64\nro.zygote=zygote64_32" >> /system/build.prop >> /system/vendor/build.prop
```

Mengatur izin direktori Android agar kembali seperti semula.

```
mount -o ro,remount /
```

Restart.

```
reboot
```
