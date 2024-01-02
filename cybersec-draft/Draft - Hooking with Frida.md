# Study-case: Root Detection in AndroGoat.apk

<https://github.com/satishpatnayak/AndroGoat>

Deteksi Root yang digunakan pada aplikasi AndroGoat, dituliskan seperti di bawah ini.
![Real Source](https://github.com/xchopath/www.novr.one/assets/44427665/365871b3-f4c4-4c91-a341-114410f5e56b)


```js
Java.perform(
  function () {
    // Bypass Root Detection
    let RootDetect = Java.use("owasp.sat.agoat.RootDetectionActivity");
    RootDetect["isRooted"].implementation = function() { return false; };
    RootDetect["isRooted1"].implementation = function() { return false; };
  }
);
```
