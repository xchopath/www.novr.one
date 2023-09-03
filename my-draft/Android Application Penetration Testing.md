# Android Application Penetration Testing

## Framework Classification

Pada awalnya saya bingung untuk melakukan Test pada aplikasi Mobile Android, dikarenakan terdapat banyak sekali variasi tutorial yang dipecah berdasarkan Framework yang digunakan, seperti `How to SSL Pinning in Flutter`, `Cordova`, dan lainnya. Tapi, dari situlah muncul satu pertanyaan di otak saya.

> *Bagaimana cara mendeteksi Framework yang digunakan oleh APK?*

Setelah saya baca beberapa artikel dan bertanya kepada teman-teman saya, dapatlah sebuah jawaban, yaitu kita dapat mendeteksinya dengan cara mencocokkan File-File yang terdapat di dalam APK.

### APK Extraction

Sebelum mencocokkan File-File di dalamnya, berikut ini adalah Command untuk ekstraksi APK menggunakan `apktool`.

```
apktool d SampleApp.apk
```

### Framework Common Files

Sebagai contoh kita akan gunakan List Common Files di bawah ini untuk mencocokkan Framework yang digunakan.

1. Flutter
```
lib/armeabi-v7a/libflutter.so
lib/arm64-v8a/libflutter.so
lib/x86/libflutter.so
lib/x86_64/libflutter.so
```

2. Cordova / PhoneGap
```
assets/www/
assets/www/index.html
assets/www/cordova.js
assets/www/cordova_plugins.js
```

3. React Native
```
lib/armeabi-v7a/libreactnativejni.so
lib/arm64-v8a/libreactnativejni.so
lib/x86/libreactnativejni.so
lib/x86_64/libreactnativejni.so
assets/index.android.bundle
```

Untuk mempermudah, di sini saya sudah menyediakan sebuah Script menggunakan GoLang untuk melakukan pengecekan Framework pada APK <https://www.novr.one/files/apkrecognize.go> (Script ini hasil modifikasi milik teman saya).

**Note:** Bilamana tidak ada Common Files milik Framework terkait, maka kemungkinan besar APK tersebut dibangun secara **Native** atau juga bisa diasumsikan kalau terdapat file `kotlin/` itu juga **Native**.

## Basic Preparation for Intercept

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
