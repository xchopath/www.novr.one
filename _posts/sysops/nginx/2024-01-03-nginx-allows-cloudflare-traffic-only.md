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

Create a file to store Cloudflare IP list.
```sh
vim /etc/nginx/cfwhitelist.conf
```

Use this command to get the IP list from Cloudflare.
```bash
{ curl -s "https://www.cloudflare.com/ips-v4"; curl -s "https://www.cloudflare.com/ips-v6"; } | awk '{print "set_real_ip_from "$1";"}' && echo "real_ip_header CF-Connecting-IP;" && echo "" && echo "geo \$realip_remote_addr \$cf_ipswhitelist {" && echo "\tdefault 0;" && { curl -s "https://www.cloudflare.com/ips-v4"; curl -s "https://www.cloudflare.com/ips-v6"; } | awk '{print "\t"$1" 1;"}' && echo "}"
```

![cfipwhitelist](https://github.com/xchopath/www.novr.one/assets/44427665/1fb603f0-6fe0-4b87-a92d-f47ca650047b)

<br/>

Add the configuration file to `/etc/nginx/nginx.conf`:
```sh
vim /etc/nginx/nginx.conf
```
```
include /etc/nginx/cfwhitelist.conf;
```

![cfwhitelist nginx conf](https://github.com/xchopath/www.novr.one/assets/44427665/535132fd-cbe1-437e-ae35-d1c56dd21144)


<br/>


# Implement to Server Host / Virtual Host Configuration

Configure host file / virtual host file add this configuration below.

```sh
if ($cf_ipswhitelist != 1) {
	return 403;
}
```

![cfwl](https://github.com/xchopath/www.novr.one/assets/44427665/3c2e9e24-4f24-4dab-8504-08620cdd1a11)

<br/>

Restart!
