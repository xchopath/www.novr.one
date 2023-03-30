# SSH Persistent Port Tunneling with systemd

- Create systemd service file:

```
sudo vim /etc/systemd/system/<service>-tunnel.service
```

- Edit file `/etc/systemd/system/<service>-tunnel.service`

For example, tunnel port `9200`.

```
[Unit]
Description=Persistent SSH Tunnel
After=network.target

[Service]
Restart=on-failure
RestartSec=5
ExecStart=/usr/bin/ssh -NTC -o ServerAliveInterval=60 -o ExitOnForwardFailure=yes -o StrictHostKeyChecking=no -L 9200:127.0.0.1:9200 user@192.168.1.100 -i /home/some/.ssh/some.id_rsa

[Install]
WantedBy=multi-user.target
```

- Enable & daemon reload systemd

```
sudo systemctl daemon-reload
sudo systemctl enable <service>-tunnel.service
sudo systemctl start <service>-tunnel.service
```
