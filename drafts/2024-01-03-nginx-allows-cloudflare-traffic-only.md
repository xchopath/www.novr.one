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

# Configure nginx.conf

Create a file to store Cloudflare whitelisted IP list.
```sh
vim /etc/nginx/cloudflare-whitelist.conf
```

Use this command to get the IP list from Cloudflare.
```bash
{ curl -s "https://www.cloudflare.com/ips-v4"; echo ""; curl -s "https://www.cloudflare.com/ips-v6"; } | awk '{print "set_real_ip_from "$1";"}' && echo "real_ip_header CF-Connecting-IP;" && echo "" && echo "geo \$realip_remote_addr \$cloudflareips {" && echo "\tdefault 0;" && { curl -s "https://www.cloudflare.com/ips-v4"; echo ""; curl -s "https://www.cloudflare.com/ips-v6"; } | awk '{print "\t"$1" 1;"}' && echo "}"
```

![cloudflare-whitelist.conf](https://github.com/xchopath/www.novr.one/assets/44427665/73d0a764-efdc-488a-aa4b-759be1d17fb3)

<br/>

Add the configuration file to `/etc/nginx/nginx.conf`:
```sh
vim /etc/nginx/nginx.conf
```

```
include /etc/nginx/cloudflare-whitelist.conf;
```

![nginx.conf](https://github.com/xchopath/www.novr.one/assets/44427665/34b787c6-7c69-4f54-b5a3-8e8c573418f3)

<br/>


# Implement to Server Host / Virtual Host Configuration

Configure host file / virtual host file add this configuration below.

```sh
if ($cloudflareips != 1) {
	return 403;
}
```

![host configuration](https://github.com/xchopath/www.novr.one/assets/44427665/23cb9faa-ed91-4407-8856-3e947ea401d0)


<br/>

Restart!
