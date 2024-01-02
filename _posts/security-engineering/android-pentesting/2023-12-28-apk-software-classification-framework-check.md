---
layout: post
title: "APK Software Classification (Framework Check)"
date: 2023-12-28 00:00:00 +0700
categories: "Android-Pentesting"
---

Pada awalnya saya bingung saat melakukan Test pada aplikasi Mobile Android, dikarenakan terdapat banyak sekali variasi tutorial yang dipecah berdasarkan Framework yang digunakan, seperti How to SSL Pinning in Flutter dan lainnya. Dari situlah muncul satu pertanyaan di otak saya.

> Bagaimana cara kita mengetahui Framework apa yang digunakan pada APK tersebut?

Setelah saya berdiskusi dengan beberapa teman saya, akhirnya mendapat pencerahan, kita dapat mengetahuinya dengan cara membaca Common Files yang tersimpan di dalam APK-nya.

## APK Extraction

Ekstrak APK dengan apktool
```sh
apktool d yours.apk -o path-to
```

Print semua file APK yang sudah diekstraksi
```sh
cd path-to/
find .
```

![List Tree of Files](https://github.com/xchopath/www.novr.one/assets/44427665/f42358f6-60e7-445a-af63-4aca75eafd9a)

<br/>

# Framework Classification

Untuk mengetahui APK tersebut dibangun menggunakan Framework dan Software apa, kita hanya perlu mencocokkannya dengan Common Files yang sudah dirangkum di bawah ini.

### Flutter
- `lib/armeabi-v7a/libflutter.so`
- `lib/arm64-v8a/libflutter.so`
- `lib/x86/libflutter.so`
- `lib/x86_64/libflutter.so`

### React Native
- `lib/armeabi-v7a/libreactnativejni.so`
- `lib/arm64-v8a/libreactnativejni.so`
- `lib/x86/libreactnativejni.so`
- `lib/x86_64/libreactnativejni.so`
- `assets/index.android.bundle`

### Kotlin (Java)
- `kotlin/*`
- `original/META-INF/kotlin-stdlib*`
- `kotlinx/*`

### Java
- `*/*.smali`

### Xamarin
- `lib/armeabi-v7a/libmonodroid.so`
- `lib/arm64-v8a/libmonodroid.so`
- `lib/x86/libmonodroid.so`
- `lib/x86_64/libmonodroid.so`
- `lib/armeabi-v7a/libmonosgen-2.0.so`
- `lib/arm64-v8a/libmonosgen-2.0.so`
- `lib/x86/libmonosgen-2.0.so`
- `lib/x86_64/libmonosgen-2.0.so`
- `assemblies/Sikur.Monodroid.dll`
- `assemblies/Sikur.dll`
- `assemblies/Xamarin.Mobile.dll`
- `assemblies/mscorlib.dll`

### Corona SDK
- `lib/armeabi-v7a/libcorona.so`
- `lib/arm64-v8a/libcorona.so`
- `lib/x86/libcorona.so`
- `lib/x86_64/libcorona.so`
- `assets/resource.car`

### Cordova / PhoneGap
- `assets/www/index.html`
- `assets/www/cordova.js`
- `assets/www/cordova_plugins.js`
