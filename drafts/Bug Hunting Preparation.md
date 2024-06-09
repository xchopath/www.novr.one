# Bug Hunting Preparation

Halaman ini berisi catatan pribadi, tentang hal yang repetitif, namun, saya sendiri sering lupa.

# Burp Suite Configurations

### Scope settings

Go to `Target` > `Scope settings` > `Use advanced scope control`.

- Protocol: `Any`
- Host or IP range: `^.*\.?domain\.com$`
- Port: `^[0-9]*$`
- File: `^/.*`

### Proxy settings

Go to `Proxy` > `Proxy settings`.

Request interception rules:
- Enable `And - URL - Is in target scope`.

Response interception rules:
- Enable `And - URL - Is in target scope`.

WebSocket interception rules:
- Enable `Only intercept in scope messages`.

Proxy logging history:
- Enable `Stop logging out-of-scope items`.

Miscellaneous:
- Enable `Don't send items to Proxy history or live tasks, if out of scope`.
