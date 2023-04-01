# Nginx Rate-Limit

Rate Limit is used to limit HTTP requests when a client is trying to flood requests. This can be used to withstand DDoS and Automated Scanner attacks.

## Nginx settings

### 1. limit_req_zone

`limit_req_zone` is used to limit the Shared-Memory Zone, but not the request rate. This parameter is implemented inside `http { }` in `/etc/nginx/nginx.conf`.

```
http {
  limit_req_zone $binary_remote_addr zone=mylimitname:10m rate=10r/s;
}
```

- `$binary_remote_addr` - Stores the binary form of the Client IP address. Performs a limit using the third parameter of the IP Address form.
- `zone=mylimitname:10m` - Specifies the Shared-Memory Zone that will be used to store the state of each IP address and determine how much that address can access the restricted URL.
- `rate=10r/s` - Sets the maximum request rate. Example, the Client should not exceed 10 requests per second.


### 2. limit_req

`limit_req` used to limit the request rate, this parameter is implemented inside `location { }` in the configuration of each host in the `/etc/nginx/sites-*/` or `/etc/nginx/conf.d/` directory.

```
location / {
  limit_req zone=mylimitname burst=5 nodelay;
  limit_req_status 429;
}
```

- `zone=mylimitname` - Using the Shared-Memory Zone that was created in the previous `limit_req_zone` parameter.
- `burst=5` - Limits the maximum requests that can be performed simultaneously.
- `nodelay` - Eliminate Delay Limitation when the Server makes a Response to the Client.
- `limit_req_status` - Change the HTTP Response Code when the client has hit the limit.

---

Restart nginx service.

```
sudo service nginx restart
```

**Note:** This is only useful to limit requests, to do auto-block you can use a service like `fail2ban`.
