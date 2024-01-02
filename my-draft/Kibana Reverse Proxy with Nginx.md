# Kibana Reverse Proxy with Nginx

## Kibana

Add or edit `server.basePath` in `/etc/kibana/kibana.yml`

```
server.basePath: "/kibana"
```

Restart kibana service

```
sudo systemctl restart kibana
```

## Nginx

Edit `/etc/nginx/conf.d/domain.com.conf`

```
upstream kibanahost {
    server 127.0.0.1:5601;
    keepalive 15;
}

server {
  server_name domain.com;
  listen 80;

  access_log /var/log/nginx/domain.com.access.log;
  error_log /var/log/nginx/domain.com.error.log;

  location /kibana/ {
    proxy_pass http://kibanahost/;
    proxy_redirect off;
    proxy_buffering off;
    proxy_http_version 1.1;
    proxy_set_header Connection "Keep-Alive";
    proxy_set_header Proxy-Connection "Keep-Alive";
    rewrite ^/kibana/(.*)$ /$1 break;
  }
}
```

Restart nginx service

```
sudo systemctl restart nginx
```

#### Testing

Go to `http://nginxhost/kibana`. 
