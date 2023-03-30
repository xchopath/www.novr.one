# ElasticSearch Enable Authentication with xpack

### Enable xpack module

Edit `/etc/elasticsearch/elasticsearch.yml`

```
xpack.security.enabled: true
discovery.type: single-node
```

Restart service

```
sudo systemctl restart elasticsearch
```

### Add Superuser

Add new user

```
sudo /usr/share/elasticsearch/bin/elasticsearch-users useradd <user>
```

Set a role

```
sudo /usr/share/elasticsearch/bin/elasticsearch-users roles <user> -a superuser
```
