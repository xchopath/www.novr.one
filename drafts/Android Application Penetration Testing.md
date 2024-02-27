# Android Application Penetration Testing

## Software Classification

Pada awalnya saya bingung untuk melakukan Test pada aplikasi Mobile Android, dikarenakan terdapat banyak sekali variasi tutorial yang dipecah berdasarkan Framework yang digunakan, seperti `How to SSL Pinning in Flutter` dan lainnya. Tapi, dari situ muncul satu pertanyaan di otak saya.

> *Bagaimana cara mendeteksi Framework yang digunakan oleh APK?*

Setelah saya baca beberapa artikel dan bertanya kepada teman-teman saya, dapatlah sebuah jawaban, yaitu kita dapat mendeteksinya dengan cara mencocokkan File-File yang terdapat di dalam APK.

### APK Extraction

Sebelum mencocokkan File-File di dalamnya, berikut ini adalah Command untuk ekstraksi APK menggunakan `apktool`.

```
apktool d SampleApp.apk
```

### Common Files Recognition

Sebagai contoh kita akan gunakan List Common Files di bawah ini untuk mencocokkan Framework yang digunakan.

1. Flutter
```
lib/armeabi-v7a/libflutter.so
lib/arm64-v8a/libflutter.so
lib/x86/libflutter.so
lib/x86_64/libflutter.so
```

2. React Native
```
lib/armeabi-v7a/libreactnativejni.so
lib/arm64-v8a/libreactnativejni.so
lib/x86/libreactnativejni.so
lib/x86_64/libreactnativejni.so
assets/index.android.bundle
```

Untuk mempermudah, di sini saya sudah menyediakan sebuah Script menggunakan GoLang untuk melakukan pengecekan Framework pada APK <https://www.novr.one/files/apkrecognize.go> (Script ini hasil modifikasi milik teman saya).

**Note:** Bilamana tidak ada Common Files milik Framework terkait, maka kemungkinan besar APK tersebut dibangun secara **Native** atau juga bisa diasumsikan kalau terdapat file `kotlin/` itu juga **Native**.

<br>

----------

<br>

## Basic Preparation for Intercept

Tool:
- BurpSuite
- ADB Tool
- Emulator (saya menggunakan Genymotion dengan Device Google Nexus 5)

Download Certificate milik BurpSuite dan convert.
```
curl -s --proxy "http://127.0.0.1:8080" "http://burp/cert" -o burpsuite.der && openssl x509 -inform DER -in burpsuite.der -out burpsuite.pem
cp burpsuite.pem $(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0
```

Connect Device menggunakan ADB.
```
adb tcpip 5555
adb connect 127.0.0.1:6555
adb root
```

Menambahkan Certificate `.pem` pada System Device. 
```
adb remount
adb push $(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0 /sdcard/
echo -n "mount -o rw,remount /" | adb shell
echo -n "mv /sdcard/$(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0 /system/etc/security/cacerts" | adb shell
echo -n "chmod 644 /system/etc/security/cacerts/$(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0" | adb shell
echo -n "chown root:root /system/etc/security/cacerts/$(openssl x509 -inform PEM -subject_hash_old -in burpsuite.pem | head -1).0" | adb shell
echo -n "mount -o ro,remount /" | adb shell
echo -n "reboot" | adb shell
```

<br>

----------

<br>

## Root Detection

Root Detection adalah salah satu cara untuk menentukan apakah perangkat Android telah di-"root" atau tidak. Ketika sudah dalam kondisi Root itu akan memberikan User untuk memiliki kontrol penuh atas perangkat, tapi juga dapat membuka potensi risiko keamanan. Oleh karena itu banyak aplikasi yang mengimplementasi hal ini, tapi sebagai seorang Tester biasanya kita akan kesulitan menghadapi proteksi Root ini.

<br>

### Study-case #1: Su Binary Call / Boolean-Based

Study-case yang pertama, di sini akan membahas tentang Root Detection yang memanggil file `su`.

File `su` adalah salah satu file binary yang sering terikat dengan hak akses root, biasanya mekanisme Root Detection pada aplikasi salah satunya itu mencoba mencari file `su`.

#### Reverse Engineering

Untuk Lab-nya di sini saya menggunakan aplikasi [AndroGoat](https://github.com/satishpatnayak/AndroGoat).

Tools yang harus dipersiapkan:

|     Tools    |          Kegunaan          |                 Instalasi                 |
|:------------:|----------------------------|:-----------------------------------------:|
|     Jadx     |        Decompile APK       |        `sudo apt-get install jadx`        |
| Frida Server | Pasang pada Device Android | <https://github.com/frida/frida/releases> |
|  Frida Tool  |   Memanipulasi Activity    |      `sudo pip3 install frida-tools`      |

#### Install Latest Frida Version

```
virtualenv frida
source frida/bin/activate
pip3 install frida-tools frida==16.1.4
```


----------

**Install Frida Server on Android Device**

Download https://github.com/frida/frida/releases/latest (frida-server)

![Frida Server](https://github.com/xchopath/www.novr.one/assets/44427665/b817358c-71ab-4ec6-b48d-a6c446a7ea8c)

Push frida-server
```
adb push frida-server-* /data/local
echo -n "chmod +x /data/local/frida-server-*" | adb shell
```

Jalankan Frida Server
```
adb shell
cd /data/local
./frida-server-*
```

----------

**Flow**

Jika file `su` tidak dapat ditemukan oleh aplikasi artinya aplikasi tidak terdeteksi sebagai Root dan akan berjalan dengan semestinya. Namun, jika file `su` berhasil ditemukan oleh aplikasi, maka hal ini mengindikasikan bahwa perangkat telah di-root, biasanya aplikasi akan langsung melakukan blokir terhadap penggunaan aplikasi tersebut.

**Hands-on**

Untuk melakukan Reverse Engineering, pertama-tama lakukan Decompile terlebih dahulu File APK-nya menggunakan `jadx`.

```
jadx -d ~/path-to/decompiled ~/path-to/file.apk
```

Jika sudah, masuk ke direktori hasil Decompile dan cari menggunakan Keyword `su` (secara Recursive). Kita dapat menggunakan Command `grep` (pada Linux) serta menggunakan Regex yang biasa saya gunakan di bawah ini.

```
grep -REn '(^|[^a-zA-Z0-9])su($|[^a-zA-Z0-9])' .
```

![image](https://github.com/xchopath/www.novr.one/assets/44427665/386c0f74-e4d5-42fe-98d4-9a6f3926558a)

Karena pada tulisan ini saya menggunakan **AndroGoat** maka file yang ditemukan yaitu `./sources/owasp/sat/agoat/RootDetectionActivity.java`.

Jika sudah, buka file tersebut dan langsung menuju fungsi Root Detection-nya.

![isRoot Function](https://github.com/xchopath/www.novr.one/assets/44427665/4c9afca6-05bd-4fcd-8b41-5647ebaa7c51)

Setelah kita baca-baca Source Code-nya, ditemukan dua fungsi yang digunakan untuk pengecekan Root, yaitu `isRooted` dan `isRooted1`.

#### Bypass

Pada kasus ini kita sudah mengetahui bahwa Activity pada Root Detection-nya menggunakan fungsi `isRooted` dan `isRooted1` maka kita akan melakukan Hooking terhadap Activity tersebut menggunakan Script di bawah ini.

Hooking adalah teknik yang digunakan untuk memantau dan mengintersep pemanggilan Function atau Method di dalam sebuah aplikasi mobile. Dengan Tool Frida ini kita dapat memasangkan sebuah Script (JavaScript) ke dalam aplikasi Android yang sedang berjalan, tujuan skrip tersebut biasanya digunakan untuk memantau dan memanipulasi aktifitas di dalam aplikasi secara Real-Time.

Buat Script `bypass-root.js`:
```
Java.perform(
  function () {
    let RootDetectCall = Java.use("owasp.sat.agoat.RootDetectionActivity");
    RootDetectCall["isRooted"].implementation = function() { return false; };
    RootDetectCall["isRooted1"].implementation = function() { return false; };
  }
);
```

Jalankan Hook.
```
frida -l <script> -U -f <identifier>
```

Mengetahui identifier sebuah aplikasi.
```
frida-ps -aU
```
