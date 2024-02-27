---
layout: post
title: "Decompile Compiled index.android.bundle on React Native (Hermes JavaScript Bytecode)"
date: 2024-01-16 07:00:00 +0700
categories: "Android-Pentesting"
---

![Hermes JavaScript bytecode](https://github.com/xchopath/www.novr.one/assets/44427665/e4c80d95-2943-41b1-8f75-c3e9fc250c88)

Recognize `index.android.bundle`:
```
file index.android.bundle
```

## Installation

```
sudo apt install python3-clang -y
```

```
sudo pip3 install --upgrade git+https://github.com/P1sec/hermes-dec
```

## Decompile
```
hbc-decompiler index.android.bundle <output.js>
```

<br/>

After decompile:

![image](https://github.com/xchopath/www.novr.one/assets/44427665/e6789208-4636-4490-aba0-f65931176ce9)
