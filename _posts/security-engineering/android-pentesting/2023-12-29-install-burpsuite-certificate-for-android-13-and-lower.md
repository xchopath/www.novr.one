---
layout: post
title: "Install BurpSuite Certificate for Android 13 and Lower"
date: 2023-12-29 00:00:00 +0700
categories: "Android-Pentesting"
---

![Burp Android](https://github.com/xchopath/www.novr.one/assets/44427665/a3e3cd94-a50b-4a73-b8d7-c29b1a91cfa8)

**Catatan**: Untuk versi 13 dan versi di bawahnya.

<table>
    <thead>
        <tr>
            <th>Device</th>
            <th>Tool</th>
            <th>Requirement</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td rowspan="2">Computer</td>
            <td>BurpSuite</td>
            <td></td>
        </tr>
        <tr>
            <td>ADB</td>
            <td></td>
        </tr>
        <tr>
            <td>Android<br>(Emulator / Real Device)</td>
            <td></td>
            <td>Root Needed</td>
        </tr>
    </tbody>
</table>

----------

# BurpSuite Settings

Perlu dilakukan penyesuaian pada konfigurasi BurpSuite, dikarenakan Host (Android) yang akan kita Intercept itu berbeda dengan Host milik BurpSuite. Konfigurasi yang perlu disesuaikan terletak pada `Proxy` > `Proxy settings` > `Proxy Listeners`, yang di mana `Bind address` perlu diubah menjadi `All Interfaces` agar dapat di-listen oleh Device di luar Host (milik BurpSuite) melalui mekanisme Remote Proxy.

![BurpSuite Bind Address to All Interfaces](https://github.com/xchopath/www.novr.one/assets/44427665/f98ac249-8153-41f2-8e0c-0246e8b74fbc)

Sebagai opsional, maka kita juga dapat mengubah konfigurasi `Request handling` agar `Support invisible proxy` (jika diperlukan).

![BurpSuite Invisible Proxy](https://github.com/xchopath/www.novr.one/assets/44427665/2b9c1d40-5f81-4472-88dc-6c32f7fc887e)

# BurpSuite Certificate Installation in Android System

Pemasangan sertifikat BurpSuite di Android System itu memerlukan koneksi dari Computer ke Android dan Tool yang akan kita gunakan yaitu ADB. Untuk mengkoneksikan ADB dan Android Device dapat dilakukan dengan 2 cara, yaitu:

1. via Network (TCP)
2. via USB

Untuk tutorialnya, sangat banyak di internet dan sangat bervariasi berdasarkan jenis Handphone-nya.

Hal ini akan kita lanjutkan dengan anggapan bahwa, ADB (pada komputer) dan Android sudah berhasil terkoneksi.


Download BurpSuite Certificate (Local) & Convert:
```
curl -s --proxy "http://127.0.0.1:8080" "http://burp/cert" -o burpsuite.der
openssl x509 -inform DER -in burpsuite.der -out burpsuite.pem
cp burpsuite.pem $(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0
```

Install BurpSuite Certificate to Android System via ADB:
```
adb root
adb remount
adb push $(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0 /sdcard/
echo -n "mount -o rw,remount /" | adb shell
echo -n "mv /sdcard/$(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0 /system/etc/security/cacerts" | adb shell
echo -n "chmod 644 /system/etc/security/cacerts/$(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0" | adb shell
echo -n "chown root:root /system/etc/security/cacerts/$(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0" | adb shell
echo -n "mount -o ro,remount /" | adb shell
echo -n "reboot" | adb shell
```
