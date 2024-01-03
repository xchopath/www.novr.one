---
layout: post
title: "Nginx Allows Cloudflare Traffic Only"
date: 2023-04-29 00:00:00 +0700
categories: "Nginx"
---

<br/>

Get Cloudflare IP List:
- <https://www.cloudflare.com/ips-v4>
- <https://www.cloudflare.com/ips-v6>

<br/>

Install Nginx extra module
```sh
sudo apt-get install nginx-extras
```

<br/>

## Setup http function in nginx.conf

Config maker for `/etc/nginx/nginx.conf` inside of `http { ... }`.

Bash:
```bash
{ curl -s "https://www.cloudflare.com/ips-v4"; curl -s "https://www.cloudflare.com/ips-v6"; } | awk '{print "set_real_ip_from "$1";"}' && echo "real_ip_header CF-Connecting-IP;" && echo "" && echo "geo \$realip_remote_addr \$cf_ipswhitelist {" && echo "\tdefault 0;" && { curl -s "https://www.cloudflare.com/ips-v4"; curl -s "https://www.cloudflare.com/ips-v6"; } | awk '{print "\t"$1" 1;"}' && echo "}"
```

<table>
  <td>
    <img src="https://github.com/xchopath/www.novr.one/assets/44427665/fc09cf6c-323f-44be-bd88-0c2addf969ed"/>
  </td>
  <td>
    <img src="https://github.com/xchopath/www.novr.one/assets/44427665/6cc00180-48b1-44f3-a93c-94dd52b4afb7">
  </td>
</table>

<br/>

## Enable rule in Server Host / Virtual Host

Config inside the `server { ... }` in Host Config File (For example: /etc/nginx/conf.d/default):
```sh
if ($cf_ipswhitelist != 1) {
	return 403;
}
```

Restart!
